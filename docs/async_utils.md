# Docs for async_utils.go

## Usage

The `async_utils.go` file in the `utils` package provides asynchronous utility functions to facilitate concurrent operations in Go. Here's a detailed breakdown of the available function:

## DecreaseChannelBatchFn

### Description
`DecreaseChannelBatchFn` manages batch processing in concurrent operations, ensuring that a set number of asynchronous tasks are completed before proceeding.

### Usage
```go
DecreaseChannelBatchFn(i, batchSize, batchDone, targetArray)
```

### Reference
| Parameter | Type | Description |
|-----------|------|-------------|
| i | `int` | The current index in the batch processing loop. |
| batchSize | `int` | The size of each batch to process. |
| batchDone | `chan bool` | A channel that signals the completion of an individual task in the batch. |
| targetArray | `[]string` | The array being processed in batches. |

#### Detailed Explanation
- When the index `i` is a multiple of `batchSize`, the function waits for `batchSize` number of tasks to signal completion on the `batchDone` channel before printing "Batch complete".
- If the end of the `targetArray` is reached and the remaining number of tasks is less than `batchSize`, the function waits for these remaining tasks to complete.
- This utility is particularly useful in scenarios where you need to process items in batches concurrently and ensure that each batch is fully processed before moving on to the next.
