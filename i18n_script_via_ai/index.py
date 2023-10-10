import argparse
from enum import Enum
import json
import os
import re
import time
import sys

def char_generator():
    encoding = sys.getdefaultencoding()
    char_count = 1
    generated_chars = set()

    while True:
        try:
            char = chr(char_count)
            if char not in generated_chars:
                generated_chars.add(char)
                yield char
            char_count += 1
        except UnicodeDecodeError:
            char_count += 1
            if char_count > len(encoding):
                char_count = 1




def local_deps():
  import sys
  if sys.platform == 'win32':
    sys.path.append(sys.path[0] + '\site-packages\windows')
  elif sys.platform =='linux':
    sys.path.append(sys.path[0] + '/site-packages/linux')
  elif sys.platform =='darwin':
    sys.path.append(sys.path[0] + '/site-packages/linux')
local_deps()
import openai
import pprint
pp = pprint.PrettyPrinter(indent=2, compact=False, width=1)


def print_if_dev(item,pretty=False):
  if pretty == True:
    pp.pprint(item)
  else:
    print(item)

def return_empty_type(value):
  if isinstance(value, str):
      return ""
  elif isinstance(value, list):
      return []
  elif isinstance(value, dict):
      return {}
  else:
      return value

def reverse_strings(value):
  if isinstance(value, str):
      # return value
      return value[::-1]
      return "0"
  elif isinstance(value, list):
      for i in range(len(value)):
          value[i] = reverse_strings(value[i])
      return value
  elif isinstance(value, dict):
      for key in value:
          value[key] = reverse_strings(value[key])
      return value
  else:
      return value

def assign_value_at_path(obj,path, value):
    keys = path.split(".")
    parent = keys[0]
    try:
      child = keys[1]
    except IndexError:
      child = None
    if child == None:
      if not parent.isdigit():
        obj[parent] =value
      else:
        obj.insert(int(parent),value)
    elif not parent.isdigit():
      if obj.get(parent) == None:
        if not child.isdigit():
          obj[parent] ={}
        else:
          obj[parent] = []
      assign_value_at_path(
        obj[parent],
        ".".join(keys[1:]),
        value
      )
    elif parent.isdigit():
      target = None
      int_parent = int(parent)
      try:
        target = obj[int_parent]
      except IndexError:
        target = None
      if target == None:
        if not child.isdigit():
          obj.insert(int_parent,{})
        else:

          obj.insert(int_parent,[])
      assign_value_at_path(
        obj[int_parent],
        ".".join(keys[1:]),
        value
      )


class OpenAIModelChatCompletionEnum(Enum):
  GPT_35_TURBO_16K ={
    "name":"gpt-3.5-turbo-16k",
    "max_tokens":int(16384  ),
    "sleep_time":60
  }
  GPT_35_TURBO_0301 ={
    "name":"gpt-3.5-turbo-0301",
    "max_tokens":int(4096 ),
    "sleep_time":60
  }
  GPT_35_TURBO = {
    "name":"gpt-3.5-turbo",
    "max_tokens":int(4096 ),
    "sleep_time":60

  }

class OpenAIManager():
  init= False
  client = None
  model = OpenAIModelChatCompletionEnum.GPT_35_TURBO_0301
  chunk_prompt = """ With the following Array `{}` recursively translate every value for each item in the list from English  to the  {} language
  MAKE SURE TO RETURN PROPER ARRAY ITS AN ARRAY OF JSON THE SOURCE RETURN AN ARRAY OF JSON THE SAME AMOUNT !!!!
  MAKE SURE TO RETURN THE SAME AMOUNT OF Array ITEMS THAT YOU ORIGINALLY RECEIVED !!!!!
  MAKE SURE TO TRANSLATE ALL THE VALUES TO THE DESINATION LANGUAGE !!!
  MAKE SURE TO TRANSLATE ALL THE VALUES TO THE DESINATION LANGUAGE !!!
  MAKE SURE TO TRANSLATE ALL THE VALUES TO THE DESINATION LANGUAGE !!!
  IF THE JSON SEEMS TO BE A STRING JUST RETURN A DOUBLE QUOTED STRING RESULT NO JSON !!!!
  USE DOUBLE QUOTES FOR ALL STRINGS!!!!!!
  USE DOUBLE QUOTES FOR ALL STRINGS!!!!!!
  USE DOUBLE QUOTES FOR ALL STRINGS!!!!!!
  USE DOUBLE QUOTES FOR ALL STRINGS!!!!!!
  USE DOUBLE QUOTES FOR ALL STRINGS!!!!!!
  USE DOUBLE QUOTES FOR ALL STRINGS!!!!!!
  USE DOUBLE QUOTES FOR ALL STRINGS!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  ESCAPE ALL ' AND " AS NECESSARY !!!!!!!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT PUT JSON OBJECTS IN QUOTES !!!!!
  DO NOT MODIFY THE STRUCTURE OF THE JSON FOR EVERY ELEMENT IN THE ARRAY !!!
  DO NOT MODIFY THE STRUCTURE OF THE JSON FOR EVERY ELEMENT IN THE ARRAY !!!
  DO NOT MODIFY THE STRUCTURE OF THE JSON FOR EVERY ELEMENT IN THE ARRAY !!!
  DO NOT MODIFY THE STRUCTURE OF THE JSON FOR EVERY ELEMENT IN THE ARRAY !!!
  DO NOT MODIFY THE STRUCTURE OF THE JSON FOR EVERY ELEMENT IN THE ARRAY !!!
  DO NOT MODIFY THE STRUCTURE OF THE JSON FOR EVERY ELEMENT IN THE ARRAY !!!
  I ONLY WANT THE DESTINATION LANGUAGE JSON OBJECT NO EXPLANATIONS !!!!
  IF YOU GET THE SAME ARRAY AS A PROMPT YOU MADE A MISTAKE CHECK THE RESULT ,ESPECIALLY AGAINST THE PREVIOUS RESULT, YOU ARE RETURNING CAREFULLY BEFORE RETURNING !!!
  """
  prompt ="""Translate the following words `{}` from English  to the  {} language
  MAKE SURE TO RETURN THE TRANSLATED STRING NO EXPLANATIONS!!!!!!!!!!!!!!
  THIS GOES FOR THE NAME OF THE LANGUAGE TOO!!!!!!!!!!!!!
  DONT MENTION ANYTHING BUT THE RESULT!!!!!!!!!!!!
  """
  language_codes = {
    'zh': 'Mandarin Chinese',
    'ja': 'Japanese',
    'pt': 'Portuguese',
    'es': 'Spanish',
    'hi': 'Hindi',
    'uk': 'Ukrainian',
    'ar': 'Arabic',
    'bn': 'Bengali',
    'ms': 'Malay',
    'fr': 'French',
    'de': 'German',
    'sw': 'Swahili',
    'am': 'Amharic'
}

  def __init__(self,api_key):
      self.init = True
      openai.api_key = api_key
      self.retry_ask_chatgpt = self._retry_ask_chatgpt


  _retry_ask_chatgpt = 3
  def _ask_chatgpt(self,prompt,randomness=0):
    try:
      model = self.model
      if(len(prompt) > model.value["max_tokens"]):
        model = OpenAIModelChatCompletionEnum.GPT_35_TURBO_16K
      response = openai.ChatCompletion.create(
          model=model.value["name"],
          messages=[{
            "role":"user",
            "content":prompt,
          }],
           max_tokens=model.value["max_tokens"]-len(prompt),
          temperature=randomness,
      )
      # print_if_dev(response,True)
      return response.choices[0].message.content
    except BaseException as e:
      print("Chat gpt error")
      print(self.retry_ask_chatgpt)
      print(e)
      time.sleep(model.value["sleep_time"])
      if self.retry_ask_chatgpt !=0:
        self.retry_ask_chatgpt -=1
        return self._ask_chatgpt(prompt,randomness)
      else:
        return {}
    # response = openai.Completion.create(
    #   model="text-davinci-003",
    #   prompt=prompt,
    #   temperature=randomness
    # )





  def translate_object(self,dest_lang):
    def inner(target):
      prompt = self.prompt.format(
        target,dest_lang
      )
      self.retry_ask_chatgpt = self._retry_ask_chatgpt
      my_translate = self._ask_chatgpt(prompt)
      my_translate = re.sub(r'[\\\n]', '', my_translate)
      try:
        print(my_translate)
        if isinstance(target, str):
          return my_translate
        return json.loads(my_translate)
      except json.JSONDecodeError as e:
        print(e)
        return return_empty_type(target)
    return inner


  chunk ={}
  char_generator_instance =char_generator()
  debug_info = {
    "keys":[]
  }
  chunk_stats = ""
  recreated_lang_obj ={}
  system_error = 0
  system_retry_error = 0

  def retry_predicate_till_success(self,predicate,retry=3):
    result = predicate(list(self.chunk.values()))
    if len(result) != len(self.debug_info["keys"]) and retry!=0:
      self.system_retry_error += 1
      # debug
      print("retry errors")
      print(self.system_retry_error)
      print(result)
      print("result lenght is "+str(len(result)))
      print("debug keys length is "+str(len(self.debug_info["keys"])))
      #
      result = self.retry_predicate_till_success(predicate,retry-1)
    return result

  def call_predicate_with_chunk(self,predicate=reverse_strings):
    self.print_chunk_stats()
    result =self.retry_predicate_till_success(predicate)
    if len(result) == len(self.debug_info["keys"]):
      for key_path,result_item in zip(self.debug_info["keys"],result):
        assign_value_at_path(
          self.recreated_lang_obj,
          key_path,result_item
        )
    else:
      self.system_error += 1
      None
      # raise RuntimeError("""
      #   A system error occured,
      #   the chunk result does not match the keyPaths length,
      #   as the program iterates through your source object it grabs the keyPath the value,
      #   if the lengths dont match, then the program can update the object with incorrect keypaths
      #   and that spells disaster for your application code
      #   we will stop because there are no further means
      #   with which this command line program can try to fix the error
      #   contact the creator of the software with your source JSON file
      # """)


  def print_chunk_stats(self):
      self.chunk_stats += "\n" + str(len(str(self.chunk)))
      self.chunk_stats += "\n" + str(self.debug_info["keys"])
      if len(self.debug_info["keys"]) == 1:
        self.chunk_stats += "\n"+str(list(self.chunk.values())[0])


  def update_chunk(self, limit, key, value, string_len=0,predicate=reverse_strings,preserve=False):
      objKey= key[1:]
      if preserve == True:
        assign_value_at_path(
          self.recreated_lang_obj,
          objKey,value
        )
      else:
        chunk_len =len(str(self.chunk)) + string_len
        if  chunk_len >= limit -5:
          self.call_predicate_with_chunk(predicate)
          self.chunk,self.char_generator_instance = {},char_generator()
          self.debug_info["keys"] = []
        self.chunk[next(self.char_generator_instance)] = value
        self.debug_info["keys"].append(objKey)

  def transform_object_via_chunk(
    self,
    source,
    target,
    predicate=reverse_strings,
    limit=0,
    preserve=True,
    keyPath=""
  ):
    dest = return_empty_type(source)
    if isinstance(source, dict):
      for key, value in source.items():
        newKeyPath  = keyPath+"."+key
        if target.get(key,None) != None and preserve == True:
          # dest[key] = target[key]
          self.update_chunk(limit=None,key= newKeyPath,value= target[key],preserve=True)
        else:
          string_len = len(str(value))
          if string_len > limit:
            dest[key] = self.transform_object_via_chunk(
              value,
              target.get(key,return_empty_type(value)),
              predicate,
              limit,
              preserve,
              keyPath=newKeyPath
            )
          else:
            self.update_chunk(limit, newKeyPath, value, string_len,predicate)

            # dest[key]=predicate(value)
            dest[key]=value
      return dest
    elif isinstance(source, list):
        for key, value in enumerate(source):
          newKeyPath  = keyPath+"."+str(key)
          try:
            new_target = target[key]
          except BaseException:
            new_target = None
          if new_target != None and preserve == True:
            # result = dest[key]
            self.update_chunk(limit=None, key=newKeyPath, value=new_target,preserve=True)
          else:
            string_len = len(str(value))
            if string_len > limit:
              try:
                new_target = target[key]
              except IndexError:
                new_target = return_empty_type(value)
              result = self.transform_object_via_chunk(
                value,
                new_target,
                predicate,
                limit,
                preserve,
                keyPath=newKeyPath
              )
            else:
              self.update_chunk(limit, newKeyPath, value, string_len,predicate)
              # result = predicate(value)
              result = value
          dest.append(result)
        return dest
    elif isinstance(source, str):

        if len(source) > (limit*4):
          dest = value ="Could not process string value too big"
        else:
          value = source
          # dest = predicate(source)
          dest = source
        print(keyPath)
        self.update_chunk(limit, keyPath, value, len(source),predicate)
        return dest


  def transform_object_via_string(
    self,
    source,
    target,
    predicate=reverse_strings,
    limit=0,
    preserve=True,
    keyPath=""
  ):
    dest = return_empty_type(source)
    if isinstance(source, dict):
      for key, value in source.items():
        newKeyPath  = keyPath+"."+key
        target = target if target else {}
        if target.get(key,None) != None and preserve == True:
          dest[key] = target[key]
        else:
          dest[key] = self.transform_object_via_string(
            value,
            target.get(key,return_empty_type(value)),
            predicate,
            limit,
            preserve,
            keyPath=newKeyPath
          )
      return dest
    elif isinstance(source, list):
        for key, value in enumerate(source):
          newKeyPath  = keyPath+"."+str(key)
          try:
            new_target = target[key]
          except BaseException:
            new_target = None
          if new_target != None and preserve == True:
            result = dest[key]
          else:
            result = self.transform_object_via_string(
              value,
              new_target,
              predicate,
              limit,
              preserve,
              keyPath=newKeyPath
            )
          dest.append(result)
        return dest
    elif isinstance(source, str):

        if len(source) > (limit*4):
          dest = value ="Could not process string value too big"
        else:
          dest = predicate(source)
        return dest


  def update_translations(self,dev_obj,my_type):
    lang_codes = dev_obj.get("lang_codes")
    source_file = dev_obj.get("source_file")
    dest_file = dev_obj.get("dest_file")
    abs_path_source_file = dev_obj.get("abs_path_source_file")
    for x in lang_codes:
      x = re.sub(r"\s", "", x)
      full_lang_name = self.language_codes.get(x)
      self.chunk_stats += "\n for the {} language \n".format(full_lang_name)
      with open(abs_path_source_file,encoding="utf-8") as f:
        source_lang  = json.load(f)

        abs_path_dest_file = os.path.join(os.getcwd(),args.location,dest_file.replace("{}",x))
        if not os.path.exists(abs_path_dest_file):
          with open(abs_path_dest_file, 'w') as e:
            e.write("{}")
        with open(abs_path_dest_file,encoding="utf-8") as g:
          dest_lang  = json.load(g)
          new_lang =self.transform_object_via_string(
            source=source_lang,
            target=dest_lang,
            limit=self.model.value["max_tokens"] //4 ,
            preserve=True,
            predicate=self.translate_object(full_lang_name)
          )
          g.close()
          if my_type == "chunk":
            self.call_predicate_with_chunk(
              self.translate_object(full_lang_name)
            )

        self.chunk_stats += "\n system_errors\n"+ str(self.system_error)
        self.chunk_stats += "\n system_retry_errors\n"+ str(self.system_retry_error)
        with open(abs_path_dest_file,"w",encoding="utf-8") as h:
            if my_type == "chunk":
              print(json.dumps(self.recreated_lang_obj,indent=2) , file=h)
            elif my_type == "string":
              print(json.dumps(new_lang,indent=2) , file=h)

            f.close()
            h.close()


    with open("chunk_stats.txt","w",encoding="utf-8") as i:
        self.chunk_stats += "\n All languages\n"
        self.chunk_stats += "\n system_errors\n"+ str(self.system_error)
        self.chunk_stats += "\n system_retry_errors\n"+ str(self.system_retry_error)
        print(self.chunk_stats , file=i)
        i.close()




if __name__ == "__main__":
    parser = argparse.ArgumentParser(
                        prog='Translation Script',
                        description='translates angular 18n script',
                        epilog='Text at the bottom of help')
    parser.add_argument('-l','--location')
    parser.add_argument('-s','--source-file')
    parser.add_argument('-d','--dest-file',default="{}.json")
    parser.add_argument('-c','--lang-codes')
    args = parser.parse_args()
    abs_path_source_file = os.path.join(os.getcwd(),args.location,args.source_file)

    lang_codes = args.lang_codes.split(",")
    params= {
        "lang_codes":lang_codes,
        "source_file":args.source_file,
        "dest_file":args.dest_file,
        "abs_path_source_file":abs_path_source_file
    }
    mngr = OpenAIManager(os.environ.get("OPENAI_API_KEY_0"))
    mngr.update_translations(params,"string")

