// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'template_riverpod_provider.dart';

class WMLTemplateWidget extends ConsumerStatefulWidget {
  WMLTemplateWidget({super.key,required this.pageWidget});

   Widget pageWidget;
  @override
  ConsumerState<WMLTemplateWidget> createState() =>
      _WMLTemplateState();
}

class _WMLTemplateState extends ConsumerState<WMLTemplateWidget> {
  @override
  Widget build(BuildContext context) {

    final rp = ref.watch(WMLTemplateRiverpodProvider);

    return SafeArea(
        child:Scaffold(
          appBar:AppBar(),
        body: widget.pageWidget,
        ));
  }


}
