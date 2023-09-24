// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class WMLTemplateRiverpodProviderValue {

}

var WMLTemplateRiverpodProviderInstance =
    WMLTemplateRiverpodProviderValue();

class WMLTemplateRiverpodNotifier  extends Notifier<WMLTemplateRiverpodProviderValue>{


  @override
  WMLTemplateRiverpodProviderValue build() {
    return WMLTemplateRiverpodProviderInstance;
  }
}

final WMLTemplateRiverpodProvider = NotifierProvider<
    WMLTemplateRiverpodNotifier,
    WMLTemplateRiverpodProviderValue>(() {
  return WMLTemplateRiverpodNotifier();
});

