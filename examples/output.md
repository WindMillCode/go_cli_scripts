
# FileName: \flutter_translate.dart 

library flutter_translate;

export 'src/utils/utils.dart';
export 'src/services/localization.dart';
export 'src/delegates/localization_delegate.dart';
export 'src/widgets/localization_provider.dart';
export 'src/widgets/localized_app.dart';
export 'src/widgets/localized_app_state.dart';
export 'src/interfaces/translate_preferences.dart';
export 'src/utils/device_locale.dart';
# FileName: \src\constants\constants.dart 

class Constants
{
    static const String assetManifestFilename = 'AssetManifest.json';

    static const String localizedAssetsPath = 'assets/i18n';

    static const String pluralZero = 'zero';
    static const String pluralOne = 'one';
    static const String pluralTwo = 'two';
    static const String pluralFew = 'few';
    static const String pluralMany = 'many';
    static const String pluralOther = 'other';

    static const String pluralValueArg = '{{value}}';
}
# FileName: \src\delegates\localization_delegate.dart 

import 'package:flutter/widgets.dart';
import 'package:intl/intl.dart';
import 'package:flutter_translate/flutter_translate.dart';
import 'package:flutter_translate/src/constants/constants.dart';
import 'package:flutter_translate/src/services/locale_service.dart';
import 'package:flutter_translate/src/validators/configuration_validator.dart';

class LocalizationDelegate extends LocalizationsDelegate<Localization>
{
    Locale? _currentLocale;

    final Locale fallbackLocale;

    final List<Locale> supportedLocales;

    final Map<Locale, String> supportedLocalesMap;

    final ITranslatePreferences? preferences;

    LocaleChangedCallback? onLocaleChanged;

    Locale get currentLocale => _currentLocale!;

    LocalizationDelegate._(this.fallbackLocale, this.supportedLocales, this.supportedLocalesMap, this.preferences);

    Future changeLocale(Locale newLocale) async
    {
        var isInitializing = _currentLocale == null;

        var locale = LocaleService.findLocale(newLocale, supportedLocales) ?? fallbackLocale;

        if(_currentLocale == locale) return;

        var localizedContent = await LocaleService.getLocaleContent(locale, supportedLocalesMap);

        Localization.load(localizedContent);

        _currentLocale = locale;

        Intl.defaultLocale = _currentLocale?.languageCode;

        if(onLocaleChanged != null)
        {
           await onLocaleChanged!(locale);
        }

        if(!isInitializing && preferences != null)
        {
           await preferences!.savePreferredLocale(locale);
        }
    }

    @override
    Future<Localization> load(Locale newLocale) async
    {
        if(currentLocale != newLocale)
        {
            await changeLocale(newLocale);
        }

        return Localization.instance;
    }


    @override
    bool isSupported(Locale? locale) => locale != null;

    @override
    bool shouldReload(LocalizationsDelegate<Localization> old) => true;

    static Future<LocalizationDelegate> create({
        required String fallbackLocale,
        required List<String> supportedLocales,
        String basePath = Constants.localizedAssetsPath,
        ITranslatePreferences? preferences}) async
    {
        WidgetsFlutterBinding.ensureInitialized();

        var fallback = localeFromString(fallbackLocale);
        var localesMap = await LocaleService.getLocalesMap(supportedLocales, basePath);
        var locales = localesMap.keys.toList();

        ConfigurationValidator.validate(fallback, locales);

        var delegate = LocalizationDelegate._(fallback, locales, localesMap, preferences);

        if(!await delegate._loadPreferences())
        {
            await delegate._loadDeviceLocale();
        }

        return delegate;
    }

    Future<bool> _loadPreferences() async
    {
        if(preferences == null) return false;

        Locale? locale;

        try
        {
            locale = await preferences!.getPreferredLocale();
        }
        catch(e)
        {
            return false;
        }

        if(locale != null)
        {
            await changeLocale(locale);
            return true;
        }

        return false;
    }

    Future _loadDeviceLocale() async
    {
        try
        {
            var locale = getCurrentLocale();

            if(locale != null)
            {
                await changeLocale(locale);
            }
        }
        catch(e)
        {
            await changeLocale(fallbackLocale);
        }
    }
}

# FileName: \src\interfaces\translate_preferences.dart 

import 'dart:ui';

abstract class ITranslatePreferences
{
    Future savePreferredLocale(Locale locale);

    Future<Locale?> getPreferredLocale();
}

# FileName: \src\services\locale_file_service.dart 

import 'dart:convert';
import 'package:flutter/services.dart';
import 'package:flutter_translate/src/constants/constants.dart';

class LocaleFileService
{
    static Future<Map<String, String>> getLocaleFiles(List<String> locales, String basePath) async
    {
        var localizedFiles = await _getAllLocaleFiles(basePath);

        final files = new Map<String, String>();

        for(final language in locales.toSet())
        {
            var file = _findLocaleFile(language, localizedFiles, basePath);

            files[language] = file;
        }

        return files;
    }

    static Future<String?> getLocaleContent(String file) async
    {
        final ByteData? data = await rootBundle.load(file);

        if (data == null) return null;
        
        return utf8.decode(data.buffer.asUint8List());
    }

    static Future<List<String>> _getAllLocaleFiles(String basePath) async
    {
        final manifest = await rootBundle.loadString(Constants.assetManifestFilename);

        Map<String, dynamic> map = jsonDecode(manifest);

        var separator = basePath.endsWith('/') ? '' : '/';

        return map.keys.where((x) => x.startsWith('$basePath$separator')).toList();
    }

    static String _findLocaleFile(String languageCode, List<String> localizedFiles, String basePath)
    {
        var file = _getFilepath(languageCode, basePath);

        if(!localizedFiles.contains(file))
        {
            if(languageCode.contains('_'))
            {
                file = _getFilepath(languageCode.split('_').first, basePath);
            }
        }

        if(file == null)
        {
            throw new Exception('The asset file for the language "$languageCode" was not found.');
        }

        return file;
    }

    static String? _getFilepath(String languageCode, String basePath)
    {
        var separator = basePath.endsWith('/') ? '' : '/';

        return '$basePath$separator$languageCode.json';
    }
}

# FileName: \src\services\locale_service.dart 

import 'dart:convert';
import 'dart:ui';
import 'package:flutter_translate/flutter_translate.dart';

import 'locale_file_service.dart';

class LocaleService
{
    static Future<Map<Locale, String>> getLocalesMap(List<String> locales, String basePath) async
    {
        var files = await LocaleFileService.getLocaleFiles(locales, basePath);

        return files.map((x,y) => MapEntry(localeFromString(x), y));
    }

    static Locale? findLocale(Locale locale, List<Locale> supportedLocales)
    {
        // Not supported by all null safety versions
        Locale? existing; // = supportedLocales.firstWhereOrNull((x) => x == locale);

        for (var x in supportedLocales) 
        {
            if (x == locale) 
            {
                existing = x;
                break;
            }
        }

        if(existing == null)
        {
            // Not supported by all null safety versions
            // existing = supportedLocales.firstWhereOrNull((x) => x.languageCode == locale.languageCode);
            for (var x in supportedLocales) 
            {
                if (x.languageCode == locale.languageCode) 
                {
                    existing = x;
                    break;
                }
            }

        }

        return existing;
    }

    static Future<Map<String, dynamic>> getLocaleContent(Locale locale, Map<Locale, String> supportedLocales) async
    {
        var file = supportedLocales[locale];

        if (file == null) return {};

        var content = await LocaleFileService.getLocaleContent(file);
        
        if (content == null) return {};

        return json.decode(content);
    }


}

# FileName: \src\services\localization.dart 

import 'package:flutter_translate/src/constants/constants.dart';
import 'package:intl/intl.dart';

class Localization {
  late Map<String, dynamic> _translations;

  Localization._();

  static Localization? _instance;
  static Localization get instance =>
      _instance ?? (_instance = Localization._());

  static void load(Map<String, dynamic> translations) {
    instance._translations = translations;
  }

  String translate(String key, {Map<String, dynamic>? args}) {
    var translation = getValueAtKeyPath(key);

    if (translation != null && translation is String && args != null) {
      translation = _assignArguments(translation, args);
    }

    return translation is String ? translation : key;
  }

  String plural(String key, num value, {Map<String, dynamic>? args}) {
    var forms = getValueAtKeyPath(key);

    if (forms is! Map) {
      return key; // Return the key if the expected plural forms are not found.
    }

    return Intl.plural(
      value,
      zero: _putArgs(forms[Constants.pluralZero], value, args: args),
      one: _putArgs(forms[Constants.pluralOne], value, args: args),
      two: _putArgs(forms[Constants.pluralTwo], value, args: args),
      few: _putArgs(forms[Constants.pluralFew], value, args: args),
      many: _putArgs(forms[Constants.pluralMany], value, args: args),
      other: _putArgs(forms[Constants.pluralOther], value, args: args) ??
          '$key.${Constants.pluralOther}',
    );
  }

  String? _putArgs(String? template, num value, {Map<String, dynamic>? args}) {
    if (template == null) {
      return null;
    }

    template = template.replaceAll(Constants.pluralValueArg, value.toString());

    if (args == null) {
      return template;
    }

    for (String k in args.keys) {
      template = template!.replaceAll("{$k}", args[k].toString());
    }

    return template;
  }

  String _assignArguments(String value, Map<String, dynamic> args) {
    for (final key in args.keys) {
      value = value.replaceAll('{$key}', '${args[key]}');
    }

    return value;
  }

  String? _getTranslation(String key, Map<String, dynamic> map) {
    List<String> keys = key.split('.');

    if (keys.length > 1) {
      var firstKey = keys.first;

      if (map.containsKey(firstKey) && map[firstKey] is! String) {
        return _getTranslation(
            key.substring(key.indexOf('.') + 1), map[firstKey]);
      }
    }

    return map[key];
  }

  Map<String, String> _getAllPluralForms(String key, Map<String, dynamic> map) {
    List<String> keys = key.split('.');

    if (keys.length > 1) {
      var firstKey = keys.first;

      if (map.containsKey(firstKey) && map[firstKey] is! String) {
        return _getAllPluralForms(
            key.substring(key.indexOf('.') + 1), map[firstKey]);
      }
    }

    final result = <String, String>{};

    for (String k in map[key].keys) {
      result[k] = map[key][k].toString();
    }

    return result;
  }

  dynamic getValueAtKeyPath(String keyPath) {
    List<String> parts = keyPath.split('.');
    dynamic current = _translations;
    for (String part in parts) {
      if (current is Map) {
        current = current[part];
      } else if (current is List) {
        int index = int.tryParse(part) ?? -1;
        if (index >= 0 && index < current.length) {
          current = current[index];
        } else {
          return null;
        }
      } else {
        return null;
      }
    }
    return current;
  }
}

# FileName: \src\services\localization_configuration.dart 

import 'package:flutter/widgets.dart';
import 'package:flutter_translate/flutter_translate.dart';
import 'package:flutter_translate/src/constants/constants.dart';
import 'locale_file_service.dart';

class LocalizationConfiguration
{
    Map<Locale, String>? _localizations;

    Map<Locale, String>? get localizations => _localizations;

    final Locale fallbackLocale;

    final List<Locale> supportedLocales;

    LocalizationConfiguration._(this.fallbackLocale, this.supportedLocales);

    static Future<LocalizationConfiguration> create(String fallbackLanguage, List<String> supportedLanguages, {String basePath = Constants.localizedAssetsPath}) async
    {
        var configuration = new LocalizationConfiguration._(localeFromString(fallbackLanguage), _generateSupportedLocales(supportedLanguages));

        _validateConfiguration(fallbackLanguage, supportedLanguages);

        var files = await LocaleFileService.getLocaleFiles(supportedLanguages, basePath);

        configuration._localizations = files.map((x,y) => _getLocalizedEntry(x, y));

        return configuration;
    }

    static void _validateConfiguration(String fallbackLanguage, List<String> supportedLanguages)
    {
        if(!supportedLanguages.contains(fallbackLanguage))
        {
            throw new Exception('The fallbackLanguage [$fallbackLanguage] must be present in the supportedLanguages list [${supportedLanguages.join(", ")}].');
        }
    }

    static List<Locale> _generateSupportedLocales(List<String> supportedLanguages)
    {
        return supportedLanguages.map((x) => localeFromString(x, languageCodeOnly: true)).toSet().toList();
    }

    static MapEntry<Locale, String> _getLocalizedEntry(String languageCode, String file)
    {
        Locale locale;

        if(languageCode.contains('_'))
        {
            var parts = languageCode.split('_');

            locale = new Locale(parts[0], parts[1]);
        }
        else
        {
            locale = new Locale(languageCode);
        }

        return MapEntry(locale, file);
    }
}

# FileName: \src\utils\device_locale.dart 

import 'package:flutter/widgets.dart';
import 'package:universal_io/io.dart';

/// Returns the current device locale
Locale? getCurrentLocale() 
{
    return _localeFromString(Platform.localeName);
}

/// Returns preferred device locales
List<Locale>? getPreferredLocales() 
{
    final deviceLocales = WidgetsBinding.instance.window.locales;

    return deviceLocales;
}

Locale? _localeFromString(String code) 
{
    var separator = code.contains('_') ? '_' : code.contains('-') ? '-' : null;

    if (separator != null) 
    {
        var parts = code.split(RegExp(separator));

        return Locale(parts[0], parts[1]);
    } 
    else 
    {
        return Locale(code);
    }
}

# FileName: \src\utils\utils.dart 

import 'package:flutter/widgets.dart';
import 'package:flutter_translate/flutter_translate.dart';

typedef LocaleChangedCallback = Future Function(Locale locale);

Locale localeFromString(String code, {bool languageCodeOnly = false}) {
  if (code.contains('_')) {
    var parts = code.split('_');

    return languageCodeOnly ? Locale(parts[0]) : Locale(parts[0], parts[1]);
  } else {
    return Locale(code);
  }
}

String localeToString(Locale locale) {
  return locale.countryCode != null
      ? '${locale.languageCode}_${locale.countryCode}'
      : locale.languageCode;
}

/// Translate the selected key into the currently selected locale
String translate(String key, {Map<String, dynamic>? args}) {
  return Localization.instance.translate(key, args: args);
}

/// Translate the selected key into the currently selected locale using pluralization
String translatePlural(String key, num value, {Map<String, dynamic>? args}) {
  return Localization.instance.plural(key, value, args: args);
}

/// Change the currently selected locale
Future changeLocale(BuildContext context, String? localeCode) async {
  if (localeCode != null) {
    await LocalizedApp.of(context)
        .delegate
        .changeLocale(localeFromString(localeCode));

    LocalizationProvider.of(context).state.onLocaleChanged();
  }
}

// get a list or map for more manpulateve options
dynamic getValueAtKeyPath(String keyPath) {
  return Localization.instance.getValueAtKeyPath(keyPath);
}

# FileName: \src\validators\configuration_validator.dart 

import 'dart:ui';

class ConfigurationValidator
{
    static void validate(Locale fallbackLocale, List<Locale> supportedLocales)
    {
        if(!supportedLocales.contains(fallbackLocale))
        {
            throw new Exception('The locale [$fallbackLocale] must be present in the list of supported locales [${supportedLocales.join(", ")}].');
        }
    }
}

# FileName: \src\widgets\localization_provider.dart 

import 'package:flutter/widgets.dart';
import 'package:flutter_translate/flutter_translate.dart';

class LocalizationProvider extends InheritedWidget
{
    final LocalizedAppState state;

    final Widget child;

    LocalizationProvider({Key? key, required this.child, required this.state}) : super(key: key, child: child);

    static LocalizationProvider of(BuildContext context) => (context.dependOnInheritedWidgetOfExactType<LocalizationProvider>())!;

    @override
    bool updateShouldNotify(LocalizationProvider oldWidget) => true;
}

# FileName: \src\widgets\localized_app.dart 

import 'package:flutter/widgets.dart';
import 'package:flutter_translate/flutter_translate.dart';

class LocalizedApp extends StatefulWidget
{
    final Widget child;

    final LocalizationDelegate delegate;

    LocalizedApp(this.delegate, this.child);

    LocalizedAppState createState() => LocalizedAppState();

    static LocalizedApp of(BuildContext context) => context.findAncestorWidgetOfExactType<LocalizedApp>()!;
}

# FileName: \src\widgets\localized_app_state.dart 

import 'package:flutter/widgets.dart';
import 'localized_app.dart';
import 'localization_provider.dart';

class LocalizedAppState extends State<LocalizedApp>
{
    void onLocaleChanged() => setState(() {});

    @override
    Widget build(BuildContext context) => LocalizationProvider(state: this, child: widget.child);
}
