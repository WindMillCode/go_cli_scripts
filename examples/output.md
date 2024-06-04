
# FileName: \api\README.md 

## HTTP API

In progress...

The API methods available on `YOUR_ROOT_PATH` + `path` option from config.

So, the base path should be like `http://127.0.0.1:9000/` or `http://127.0.0.1:9000/myapp/` if `path` option was set to `/myapp`.

Endpoints:

- GET `/` - return a JSON to test the server.

This group of methods uses `:key` option from config:

- GET `/:key/id` - return a new user id. required `:key` from config.
- GET `/:key/peers` - return an array of all connected users. required `:key` from config. **IMPORTANT:** You should set `allow_discovery` to `true` in config to enable this method. It disabled by default.

# FileName: \api\index.ts 

import cors, { CorsOptions } from "cors";
import express from "express";
import publicContent from "../../app.json";
import PublicApi from "./v1/public/index.ts";
import type { IConfig } from "../config/index.ts";
import type { IRealm } from "../models/realm.ts";

export const Api = ({
	config,
	realm,
	corsOptions,
}: {
	config: IConfig;
	realm: IRealm;
	corsOptions: CorsOptions;
}): express.Router => {
	const app = express.Router();

	app.use(cors(corsOptions));

	app.get("/", (_, res) => {
		res.send(publicContent);
	});

	app.use("/:key", PublicApi({ config, realm }));

	return app;
};

# FileName: \api\v1\public\index.ts 

import express from "express";
import type { IConfig } from "../../../config/index.ts";
import type { IRealm } from "../../../models/realm.ts";

export default ({
	config,
	realm,
}: {
	config: IConfig;
	realm: IRealm;
}): express.Router => {
	const app = express.Router();

	// Retrieve guaranteed random ID.
	app.get("/id", (_, res: express.Response) => {
		res.contentType("html");
		res.send(realm.generateClientId(config.generateClientId));
	});

	// Get a list of all peers for a key, enabled by the `allowDiscovery` flag.
	app.get("/peers", (_, res: express.Response) => {
		if (config.allow_discovery) {
			const clientsIds = realm.getClientsIds();

			return res.send(clientsIds);
		}

		return res.sendStatus(401);
	});

	return app;
};

# FileName: \config\index.ts 

import type { WebSocketServer, ServerOptions } from "ws";
import type { CorsOptions } from "cors";

export interface IConfig {
	readonly host: string;
	readonly port: number;
	readonly expire_timeout: number;
	readonly alive_timeout: number;
	readonly key: string;
	readonly path: string;
	readonly concurrent_limit: number;
	readonly allow_discovery: boolean;
	readonly proxied: boolean | string;
	readonly cleanup_out_msgs: number;
	readonly ssl?: {
		key: string;
		cert: string;
	};
	readonly generateClientId?: () => string;
	readonly createWebSocketServer?: (options: ServerOptions) => WebSocketServer;
	readonly corsOptions: CorsOptions;
	readonly server_type: "websocket" | "socketio";
}

const defaultConfig: IConfig = {
	host: "::",
	port: 9000,
	expire_timeout: 5000,
	alive_timeout: 90000,
	key: "peerjs",
	path: "/",
	concurrent_limit: 5000,
	allow_discovery: false,
	proxied: false,
	cleanup_out_msgs: 1000,
	corsOptions: { origin: true },
	server_type: "websocket",
};

export default defaultConfig;

# FileName: \enums.ts 

export enum Errors {
	INVALID_KEY = "Invalid key provided",
	INVALID_TOKEN = "Invalid token provided",
	INVALID_WS_PARAMETERS = "No id, token, or key supplied to websocket server",
	CONNECTION_LIMIT_EXCEED = "Server has reached its concurrent user limit",
}

export enum MessageType {
	OPEN = "OPEN",
	LEAVE = "LEAVE",
	CANDIDATE = "CANDIDATE",
	OFFER = "OFFER",
	ANSWER = "ANSWER",
	EXPIRE = "EXPIRE",
	HEARTBEAT = "HEARTBEAT",
	ID_TAKEN = "ID-TAKEN",
	ERROR = "ERROR",
}

# FileName: \index.ts 

import express, { type Express } from "express";
import http from "node:http";
import https from "node:https";

import type { IConfig } from "./config/index.ts";
import defaultConfig from "./config/index.ts";
import type { PeerServerEvents } from "./instance.ts";
import { createInstance } from "./instance.ts";
import type { IClient } from "./models/client.ts";
import type { IMessage } from "./models/message.ts";

export type { MessageType } from "./enums.ts";
export type { IConfig, PeerServerEvents, IClient, IMessage };

function ExpressPeerServer(
	server: https.Server | http.Server,
	options?: Partial<IConfig>,
) {
	const app = express();

	const newOptions: IConfig = {
		...defaultConfig,
		...options,
	};

	if (newOptions.proxied) {
		app.set(
			"trust proxy",
			newOptions.proxied === "false" ? false : !!newOptions.proxied,
		);
	}

	app.on("mount", () => {
		// eslint-disable-next-line @typescript-eslint/no-unnecessary-condition
		if (!server) {
			throw new Error(
				"Server is not passed to constructor - " + "can't start PeerServer",
			);
		}

		createInstance({ app, server, options: newOptions });
	});

	return app as Express & PeerServerEvents;
}

function PeerServer(
	options: Partial<IConfig> = {},
	callback?: (server: https.Server | http.Server) => void,
) {
	const app = express();

	let newOptions: IConfig = {
		...defaultConfig,
		...options,
	};

	const port = newOptions.port;
	const host = newOptions.host;

	let server: https.Server | http.Server;

	const { ssl, ...restOptions } = newOptions;
	if (ssl && Object.keys(ssl).length) {
		server = https.createServer(ssl, app);

		newOptions = restOptions;
	} else {
		server = http.createServer(app);
	}

	const peerjs = ExpressPeerServer(server, newOptions);
	app.use(peerjs);

	server.listen(port, host, () => callback?.(server));

	return peerjs;
}

export { ExpressPeerServer, PeerServer };

# FileName: \instance.ts 

import type express from "express";
import type { Server as HttpServer } from "node:http";
import type { Server as HttpsServer } from "node:https";
import path from "node:path";
import type { IRealm } from "./models/realm.ts";
import { Realm } from "./models/realm.ts";
import { CheckBrokenConnections } from "./services/checkBrokenConnections/index.ts";
import type { IMessagesExpire } from "./services/messagesExpire/index.ts";
import { MessagesExpire } from "./services/messagesExpire/index.ts";
import type { IWebSocketServer } from "./services/webSocketServer/index.ts";
import { WebSocketServer } from "./services/webSocketServer/index.ts";
import { MessageHandler } from "./messageHandler/index.ts";
import { Api } from "./api/index.ts";
import type { IClient } from "./models/client.ts";
import type { IMessage } from "./models/message.ts";
import type { IConfig } from "./config/index.ts";
import { ISocketIOServer,SocketIOServer } from "./services/socketioServer/index.ts";

export interface PeerServerEvents {
	on(event: "connection", listener: (client: IClient) => void): this;
	on(event: "message",listener: (client: IClient, message: IMessage) => void): this;
	// eslint-disable-next-line @typescript-eslint/unified-signatures
	on(event: "disconnect", listener: (client: IClient) => void): this;
	on(event: "error", listener: (client: Error) => void): this;
}

export const createInstance = ({
	app,
	server,
	options,
}: {
	app: express.Application;
	server: HttpServer | HttpsServer;
	options: IConfig;
}): void => {
	const config = options;
	const realm: IRealm = new Realm();
	const messageHandler = new MessageHandler(realm);

	const api = Api({ config, realm, corsOptions: options.corsOptions });
	const messagesExpire: IMessagesExpire = new MessagesExpire({
		realm,
		config,
		messageHandler,
	});
	const checkBrokenConnections = new CheckBrokenConnections({
		realm,
		config,
		onClose: (client) => {
			app.emit("disconnect", client);
		},
	});

	app.use(options.path, api);

	//use mountpath for socket server
	const customConfig = {
		...config,
		path: path.posix.join(app.path(), options.path, "/"),
	};

	let serverInstance: IWebSocketServer | ISocketIOServer;

  if (config.server_type === "socketio") {
    serverInstance = new SocketIOServer({
      server,
      realm,
      config: customConfig,
    });

		serverInstance.on("connect", (socket) => {
			const clientId = socket.id; // Using socket.id as the client ID
			const messageQueue = realm.getMessageQueueById(clientId);

			if (messageQueue) {
				let message;

				while ((message = messageQueue.readMessage())) {
					// Assuming messageHandler.handle can accept a socket object
					messageHandler.handle(socket, message);
				}
				realm.clearMessageQueue(clientId);
			}


		});

		// Listen for custom "message" event
		serverInstance.on("message", (client: IClient, message: IMessage) => {
			app.emit("message", client, message);
			messageHandler.handle(client, message);
		});

		// Listen for the "disconnect" event
		serverInstance.on("disconnect", (client: IClient) => {
			app.emit("disconnect", client);
		});

		// Listen for possible "error" events on the socket
		serverInstance.on("error", (error) => {
			app.emit("error", error);
		});
  } else {
    let serverInstance = new WebSocketServer({
      server,
      realm,
      config: customConfig,
    });

		serverInstance.on("connection", (client: IClient) => {
			const messageQueue = realm.getMessageQueueById(client.getId());

			if (messageQueue) {
				let message: IMessage | undefined;

				while ((message = messageQueue.readMessage())) {
					messageHandler.handle(client, message);
				}
				realm.clearMessageQueue(client.getId());
			}

			app.emit("connection", client);
		});

		serverInstance.on("message", (client: IClient, message: IMessage) => {
			app.emit("message", client, message);
			messageHandler.handle(client, message);
		});

		serverInstance.on("close", (client: IClient) => {
			app.emit("disconnect", client);
		});

		serverInstance.on("error", (error: Error) => {
			app.emit("error", error);
		});
  }




	messagesExpire.startMessagesExpiration();
	checkBrokenConnections.start();
};

# FileName: \messageHandler\handler.ts 

import type { IClient } from "../models/client.ts";
import type { IMessage } from "../models/message.ts";

export type Handler = (
	client: IClient | undefined,
	message: IMessage,
) => boolean;

# FileName: \messageHandler\handlers\heartbeat\index.ts 

import type { IClient } from "../../../models/client.ts";

export const HeartbeatHandler = (client: IClient | undefined): boolean => {
	if (client) {
		const nowTime = new Date().getTime();
		client.setLastPing(nowTime);
	}

	return true;
};

# FileName: \messageHandler\handlers\index.ts 

export { HeartbeatHandler } from "./heartbeat/index.ts";
export { TransmissionHandler } from "./transmission/index.ts";

# FileName: \messageHandler\handlers\transmission\index.ts 

import { MessageType } from "../../../enums.ts";
import type { IClient } from "../../../models/client.ts";
import type { IMessage } from "../../../models/message.ts";
import type { IRealm } from "../../../models/realm.ts";

export const TransmissionHandler = ({
	realm,
}: {
	realm: IRealm;
}): ((client: IClient | undefined, message: IMessage) => boolean) => {
	const handle = (client: IClient | undefined, message: IMessage) => {
		const type = message.type;
		const srcId = message.src;
		const dstId = message.dst;

		const destinationClient = realm.getClientById(dstId);

		// User is connected!
		if (destinationClient) {
			const socket = destinationClient.getSocket();
			try {
				if (socket) {
					const data = JSON.stringify(message);

					socket.send(data);
				} else {
					// Neither socket no res available. Peer dead?
					throw new Error("Peer dead");
				}
			} catch (e) {
				// This happens when a peer disconnects without closing connections and
				// the associated WebSocket has not closed.
				// Tell other side to stop trying.
				if (socket) {
					socket.close();
				} else {
					realm.removeClientById(destinationClient.getId());
				}

				handle(client, {
					type: MessageType.LEAVE,
					src: dstId,
					dst: srcId,
				});
			}
		} else {
			// Wait for this client to connect/reconnect (XHR) for important
			// messages.
			const ignoredTypes = [MessageType.LEAVE, MessageType.EXPIRE];

			if (!ignoredTypes.includes(type) && dstId) {
				realm.addMessageToQueue(dstId, message);
			} else if (type === MessageType.LEAVE && !dstId) {
				realm.removeClientById(srcId);
			} else {
				// Unavailable destination specified with message LEAVE or EXPIRE
				// Ignore
			}
		}

		return true;
	};

	return handle;
};

# FileName: \messageHandler\handlersRegistry.ts 

import type { MessageType } from "../enums.ts";
import type { IClient } from "../models/client.ts";
import type { IMessage } from "../models/message.ts";
import type { Handler } from "./handler.ts";

export interface IHandlersRegistry {
	registerHandler(messageType: MessageType, handler: Handler): void;
	handle(client: IClient | undefined, message: IMessage): boolean;
}

export class HandlersRegistry implements IHandlersRegistry {
	private readonly handlers = new Map<MessageType, Handler>();

	public registerHandler(messageType: MessageType, handler: Handler): void {
		if (this.handlers.has(messageType)) return;

		this.handlers.set(messageType, handler);
	}

	public handle(client: IClient | undefined, message: IMessage): boolean {
		const { type } = message;

		const handler = this.handlers.get(type);

		if (!handler) return false;

		return handler(client, message);
	}
}

# FileName: \messageHandler\index.ts 

import { MessageType } from "../enums.ts";
import { HeartbeatHandler, TransmissionHandler } from "./handlers/index.ts";
import type { IHandlersRegistry } from "./handlersRegistry.ts";
import { HandlersRegistry } from "./handlersRegistry.ts";
import type { IClient } from "../models/client.ts";
import type { IMessage } from "../models/message.ts";
import type { IRealm } from "../models/realm.ts";
import type { Handler } from "./handler.ts";

export interface IMessageHandler {
	handle(client: IClient | undefined, message: IMessage): boolean;
}

export class MessageHandler implements IMessageHandler {
	constructor(
		realm: IRealm,
		private readonly handlersRegistry: IHandlersRegistry = new HandlersRegistry(),
	) {
		const transmissionHandler: Handler = TransmissionHandler({ realm });
		const heartbeatHandler: Handler = HeartbeatHandler;

		const handleTransmission: Handler = (
			client: IClient | undefined,
			{ type, src, dst, payload }: IMessage,
		): boolean => {
			return transmissionHandler(client, {
				type,
				src,
				dst,
				payload,
			});
		};

		const handleHeartbeat = (client: IClient | undefined, message: IMessage) =>
			heartbeatHandler(client, message);

		this.handlersRegistry.registerHandler(
			MessageType.HEARTBEAT,
			handleHeartbeat,
		);
		this.handlersRegistry.registerHandler(
			MessageType.OFFER,
			handleTransmission,
		);
		this.handlersRegistry.registerHandler(
			MessageType.ANSWER,
			handleTransmission,
		);
		this.handlersRegistry.registerHandler(
			MessageType.CANDIDATE,
			handleTransmission,
		);
		this.handlersRegistry.registerHandler(
			MessageType.LEAVE,
			handleTransmission,
		);
		this.handlersRegistry.registerHandler(
			MessageType.EXPIRE,
			handleTransmission,
		);
	}

	public handle(client: IClient | undefined, message: IMessage): boolean {
		return this.handlersRegistry.handle(client, message);
	}
}

# FileName: \models\client.ts 

import type WebSocket from "ws";

export interface IClient {
	getId(): string;

	getToken(): string;

	getSocket(): WebSocket | null;

	setSocket(socket: WebSocket | null): void;

	getLastPing(): number;

	setLastPing(lastPing: number): void;

	send<T>(data: T): void;
}

export class Client implements IClient {
	private readonly id: string;
	private readonly token: string;
	private socket: WebSocket | null = null;
	private lastPing: number = new Date().getTime();

	constructor({ id, token }: { id: string; token: string }) {
		this.id = id;
		this.token = token;
	}

	public getId(): string {
		return this.id;
	}

	public getToken(): string {
		return this.token;
	}

	public getSocket(): WebSocket | null {
		return this.socket;
	}

	public setSocket(socket: WebSocket | null): void {
		this.socket = socket;
	}

	public getLastPing(): number {
		return this.lastPing;
	}

	public setLastPing(lastPing: number): void {
		this.lastPing = lastPing;
	}

	public send<T>(data: T): void {
		this.socket?.send(JSON.stringify(data));
	}
}

# FileName: \models\message.ts 

import type { MessageType } from "../enums.ts";

export interface IMessage {
	readonly type: MessageType;
	readonly src: string;
	readonly dst: string;
	readonly payload?: string | undefined;
}

# FileName: \models\messageQueue.ts 

import type { IMessage } from "./message.ts";

export interface IMessageQueue {
	getLastReadAt(): number;

	addMessage(message: IMessage): void;

	readMessage(): IMessage | undefined;

	getMessages(): IMessage[];
}

export class MessageQueue implements IMessageQueue {
	private lastReadAt: number = new Date().getTime();
	private readonly messages: IMessage[] = [];

	public getLastReadAt(): number {
		return this.lastReadAt;
	}

	public addMessage(message: IMessage): void {
		this.messages.push(message);
	}

	public readMessage(): IMessage | undefined {
		if (this.messages.length > 0) {
			this.lastReadAt = new Date().getTime();
			return this.messages.shift();
		}

		return undefined;
	}

	public getMessages(): IMessage[] {
		return this.messages;
	}
}

# FileName: \models\realm.ts 

import type { IMessageQueue } from "./messageQueue.ts";
import { MessageQueue } from "./messageQueue.ts";
import { randomUUID } from "node:crypto";
import type { IClient } from "./client.ts";
import type { IMessage } from "./message.ts";

export interface IRealm {
	getClientsIds(): string[];

	getClientById(clientId: string): IClient | undefined;

	getClientsIdsWithQueue(): string[];

	setClient(client: IClient, id: string): void;

	removeClientById(id: string): boolean;

	getMessageQueueById(id: string): IMessageQueue | undefined;

	addMessageToQueue(id: string, message: IMessage): void;

	clearMessageQueue(id: string): void;

	generateClientId(generateClientId?: () => string): string;
}

export class Realm implements IRealm {
	private readonly clients = new Map<string, IClient>();
	private readonly messageQueues = new Map<string, IMessageQueue>();

	public getClientsIds(): string[] {
		return [...this.clients.keys()];
	}

	public getClientById(clientId: string): IClient | undefined {
		return this.clients.get(clientId);
	}

	public getClientsIdsWithQueue(): string[] {
		return [...this.messageQueues.keys()];
	}

	public setClient(client: IClient, id: string): void {
		this.clients.set(id, client);
	}

	public removeClientById(id: string): boolean {
		const client = this.getClientById(id);

		if (!client) return false;

		this.clients.delete(id);

		return true;
	}

	public getMessageQueueById(id: string): IMessageQueue | undefined {
		return this.messageQueues.get(id);
	}

	public addMessageToQueue(id: string, message: IMessage): void {
		if (!this.getMessageQueueById(id)) {
			this.messageQueues.set(id, new MessageQueue());
		}

		this.getMessageQueueById(id)?.addMessage(message);
	}

	public clearMessageQueue(id: string): void {
		this.messageQueues.delete(id);
	}

	public generateClientId(generateClientId?: () => string): string {
		const generateId = generateClientId ? generateClientId : randomUUID;

		let clientId = generateId();

		while (this.getClientById(clientId)) {
			clientId = generateId();
		}

		return clientId;
	}
}

# FileName: \services\checkBrokenConnections\index.ts 

import type { IConfig } from "../../config/index.ts";
import type { IClient } from "../../models/client.ts";
import type { IRealm } from "../../models/realm.ts";

const DEFAULT_CHECK_INTERVAL = 300;

type CustomConfig = Pick<IConfig, "alive_timeout">;

export class CheckBrokenConnections {
	public readonly checkInterval: number;
	private timeoutId: NodeJS.Timeout | null = null;
	private readonly realm: IRealm;
	private readonly config: CustomConfig;
	private readonly onClose?: (client: IClient) => void;

	constructor({
		realm,
		config,
		checkInterval = DEFAULT_CHECK_INTERVAL,
		onClose,
	}: {
		realm: IRealm;
		config: CustomConfig;
		checkInterval?: number;
		onClose?: (client: IClient) => void;
	}) {
		this.realm = realm;
		this.config = config;
		this.onClose = onClose;
		this.checkInterval = checkInterval;
	}

	public start(): void {
		if (this.timeoutId) {
			clearTimeout(this.timeoutId);
		}

		this.timeoutId = setTimeout(() => {
			this.checkConnections();

			this.timeoutId = null;

			this.start();
		}, this.checkInterval);
	}

	public stop(): void {
		if (this.timeoutId) {
			clearTimeout(this.timeoutId);
			this.timeoutId = null;
		}
	}

	private checkConnections(): void {
		const clientsIds = this.realm.getClientsIds();

		const now = new Date().getTime();
		const { alive_timeout: aliveTimeout } = this.config;

		for (const clientId of clientsIds) {
			const client = this.realm.getClientById(clientId);

			if (!client) continue;

			const timeSinceLastPing = now - client.getLastPing();

			if (timeSinceLastPing < aliveTimeout) continue;

			try {
				client.getSocket()?.close();
			} finally {
				this.realm.clearMessageQueue(clientId);
				this.realm.removeClientById(clientId);

				client.setSocket(null);

				this.onClose?.(client);
			}
		}
	}
}

# FileName: \services\messagesExpire\index.ts 

import { MessageType } from "../../enums.ts";
import type { IConfig } from "../../config/index.ts";
import type { IMessageHandler } from "../../messageHandler/index.ts";
import type { IRealm } from "../../models/realm.ts";

export interface IMessagesExpire {
	startMessagesExpiration(): void;
	stopMessagesExpiration(): void;
}

type CustomConfig = Pick<IConfig, "cleanup_out_msgs" | "expire_timeout">;

export class MessagesExpire implements IMessagesExpire {
	private readonly realm: IRealm;
	private readonly config: CustomConfig;
	private readonly messageHandler: IMessageHandler;

	private timeoutId: NodeJS.Timeout | null = null;

	constructor({
		realm,
		config,
		messageHandler,
	}: {
		realm: IRealm;
		config: CustomConfig;
		messageHandler: IMessageHandler;
	}) {
		this.realm = realm;
		this.config = config;
		this.messageHandler = messageHandler;
	}

	public startMessagesExpiration(): void {
		if (this.timeoutId) {
			clearTimeout(this.timeoutId);
		}

		// Clean up outstanding messages
		this.timeoutId = setTimeout(() => {
			this.pruneOutstanding();

			this.timeoutId = null;

			this.startMessagesExpiration();
		}, this.config.cleanup_out_msgs);
	}

	public stopMessagesExpiration(): void {
		if (this.timeoutId) {
			clearTimeout(this.timeoutId);
			this.timeoutId = null;
		}
	}

	private pruneOutstanding(): void {
		const destinationClientsIds = this.realm.getClientsIdsWithQueue();

		const now = new Date().getTime();
		const maxDiff = this.config.expire_timeout;

		const seen: Record<string, boolean> = {};

		for (const destinationClientId of destinationClientsIds) {
			const messageQueue = this.realm.getMessageQueueById(destinationClientId);

			if (!messageQueue) continue;

			const lastReadDiff = now - messageQueue.getLastReadAt();

			if (lastReadDiff < maxDiff) continue;

			const messages = messageQueue.getMessages();

			for (const message of messages) {
				const seenKey = `${message.src}_${message.dst}`;

				if (!seen[seenKey]) {
					this.messageHandler.handle(undefined, {
						type: MessageType.EXPIRE,
						src: message.dst,
						dst: message.src,
					});

					seen[seenKey] = true;
				}
			}

			this.realm.clearMessageQueue(destinationClientId);
		}
	}
}

# FileName: \services\socketioServer\index.ts 

import { Server as IOServer } from "socket.io";
import { EventEmitter } from "events";
import { Errors, MessageType } from "../../enums.ts";
import { IConfig } from "../../config/index.ts";
import { IClient, Client } from "../../models/client.ts";
import { IMessage } from "../../models/message.ts";
import { IRealm } from "../../models/realm.ts";
import type { Server as HttpServer } from "node:http";
import type { Server as HttpsServer } from "node:https";

export interface ISocketIOServer extends EventEmitter {
  readonly path: string;
}

type CustomConfig = Pick<IConfig, "path" | "key" | "concurrent_limit">;

const WS_PATH = "peerjs";

export class SocketIOServer extends EventEmitter implements ISocketIOServer {
  public readonly path: string;
  private readonly realm: IRealm;
  private readonly config: CustomConfig;
  public readonly socketServer: IOServer;

  constructor({
    server,
    realm,
    config,
  }: {
    server: HttpServer | HttpsServer;
    realm: IRealm;
    config: CustomConfig;
  }) {
    super();

    this.setMaxListeners(0);

    this.realm = realm;
    this.config = config;

    const path = this.config.path;
    this.path = `${path}${path.endsWith("/") ? "" : "/"}${WS_PATH}`;

    this.socketServer = new IOServer(server, {
      path: this.path,
    });

    this.socketServer.on("connect", (socket) => {
      this._onSocketConnection(socket);
    });

    this.socketServer.on("error", (error) => {
      this._onSocketError(error);
    });
  }

  private _onSocketConnection(socket: any): void {
    socket.on("error", (error: Error) => {
      this._onSocketError(error);
    });

    const { id, token, key } = socket.handshake.query;

    if (!id || !token || !key) {
      this._sendErrorAndClose(socket, Errors.INVALID_WS_PARAMETERS);
      return;
    }

    if (key !== this.config.key) {
      this._sendErrorAndClose(socket, Errors.INVALID_KEY);
      return;
    }

    const client = this.realm.getClientById(id);

    if (client) {
      if (token !== client.getToken()) {
        socket.emit(MessageType.ID_TAKEN, { msg: "ID is taken" });
        socket.disconnect();
        return;
      }

      this._configureSocket(socket, client);
      return;
    }

    this._registerClient({ socket, id, token });
  }

  private _onSocketError(error: Error): void {
    this.emit("error", error);
  }

  private _registerClient({
    socket,
    id,
    token,
  }: {
    socket: any;
    id: string;
    token: string;
  }): void {
    const clientsCount = this.realm.getClientsIds().length;

    if (clientsCount >= this.config.concurrent_limit) {
      this._sendErrorAndClose(socket, Errors.CONNECTION_LIMIT_EXCEED);
      return;
    }

    const newClient: IClient = new Client({ id, token });
    this.realm.setClient(newClient, id);
    socket.emit(MessageType.OPEN);

    this._configureSocket(socket, newClient);
  }

  private _configureSocket(socket: any, client: IClient): void {
    client.setSocket(socket);

    socket.on("disconnect", () => {
      if (client.getSocket() === socket) {
        this.realm.removeClientById(client.getId());
        this.emit("close", client);
      }
    });

    socket.on("message", (data: any) => {
      try {
        const message = JSON.parse(data) as Writable<IMessage>;
        message.src = client.getId();
        this.emit("message", client, message);
      } catch (e) {
        this.emit("error", e);
      }
    });

    this.emit("connection", client);
  }

  private _sendErrorAndClose(socket: any, msg: Errors): void {
    socket.emit(MessageType.ERROR, { msg });
    socket.disconnect();
  }
}

type Writable<T> = {
  -readonly [K in keyof T]: T[K];
};

# FileName: \services\webSocketServer\index.ts 

import { EventEmitter } from "node:events";
import type { IncomingMessage } from "node:http";
import type WebSocket from "ws";
import { Errors, MessageType } from "../../enums.ts";
import type { IClient } from "../../models/client.ts";
import { Client } from "../../models/client.ts";
import type { IConfig } from "../../config/index.ts";
import type { IRealm } from "../../models/realm.ts";
import { WebSocketServer as Server } from "ws";
import type { Server as HttpServer } from "node:http";
import type { Server as HttpsServer } from "node:https";
import { IMessage } from "../../models/message.js";

export interface IWebSocketServer extends EventEmitter {
	readonly path: string;
}

type CustomConfig = Pick<
	IConfig,
	"path" | "key" | "concurrent_limit" | "createWebSocketServer"
>;

const WS_PATH = "peerjs";

export class WebSocketServer extends EventEmitter implements IWebSocketServer {
	public readonly path: string;
	private readonly realm: IRealm;
	private readonly config: CustomConfig;
	public readonly socketServer: Server;

	constructor({
		server,
		realm,
		config,
	}: {
		server: HttpServer | HttpsServer;
		realm: IRealm;
		config: CustomConfig;
	}) {
		super();

		this.setMaxListeners(0);

		this.realm = realm;
		this.config = config;

		const path = this.config.path;
		this.path = `${path}${path.endsWith("/") ? "" : "/"}${WS_PATH}`;

		const options: WebSocket.ServerOptions = {
			path: this.path,
			server,
		};

		this.socketServer = config.createWebSocketServer
			? config.createWebSocketServer(options)
			: new Server(options);

		this.socketServer.on("connection", (socket, req) => {
			this._onSocketConnection(socket, req);
		});
		this.socketServer.on("error", (error: Error) => {
			this._onSocketError(error);
		});
	}

	private _onSocketConnection(socket: WebSocket, req: IncomingMessage): void {
		// An unhandled socket error might crash the server. Handle it first.
		socket.on("error", (error) => {
			this._onSocketError(error);
		});

		// We are only interested in the query, the base url is therefore not relevant
		const { searchParams } = new URL(req.url ?? "", "https://peerjs");
		const { id, token, key } = Object.fromEntries(searchParams.entries());

		if (!id || !token || !key) {
			this._sendErrorAndClose(socket, Errors.INVALID_WS_PARAMETERS);
			return;
		}

		if (key !== this.config.key) {
			this._sendErrorAndClose(socket, Errors.INVALID_KEY);
			return;
		}

		const client = this.realm.getClientById(id);

		if (client) {
			if (token !== client.getToken()) {
				// ID-taken, invalid token
				socket.send(
					JSON.stringify({
						type: MessageType.ID_TAKEN,
						payload: { msg: "ID is taken" },
					}),
				);

				socket.close();
				return;
			}

			this._configureWS(socket, client);
			return;
		}

		this._registerClient({ socket, id, token });
	}

	private _onSocketError(error: Error): void {
		// handle error
		this.emit("error", error);
	}

	private _registerClient({
		socket,
		id,
		token,
	}: {
		socket: WebSocket;
		id: string;
		token: string;
	}): void {
		// Check concurrent limit
		const clientsCount = this.realm.getClientsIds().length;

		if (clientsCount >= this.config.concurrent_limit) {
			this._sendErrorAndClose(socket, Errors.CONNECTION_LIMIT_EXCEED);
			return;
		}

		const newClient: IClient = new Client({ id, token });
		this.realm.setClient(newClient, id);
		socket.send(JSON.stringify({ type: MessageType.OPEN }));

		this._configureWS(socket, newClient);
	}

	private _configureWS(socket: WebSocket, client: IClient): void {
		client.setSocket(socket);

		// Cleanup after a socket closes.
		socket.on("close", () => {
			if (client.getSocket() === socket) {
				this.realm.removeClientById(client.getId());
				this.emit("close", client);
			}
		});

		// Handle messages from peers.
		socket.on("message", (data) => {
			try {
				// eslint-disable-next-line @typescript-eslint/no-base-to-string
				const message = JSON.parse(data.toString()) as Writable<IMessage>;

				message.src = client.getId();

				this.emit("message", client, message);
			} catch (e) {
				this.emit("error", e);
			}
		});

		this.emit("connection", client);
	}

	private _sendErrorAndClose(socket: WebSocket, msg: Errors): void {
		socket.send(
			JSON.stringify({
				type: MessageType.ERROR,
				payload: { msg },
			}),
		);

		socket.close();
	}
}

type Writable<T> = {
	-readonly [K in keyof T]: T[K];
};
