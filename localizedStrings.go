//////////////////////////////////////////////////////////////////////
// localizedStrings.go
//
// @usage
// 
//   1. Set json files you want to load.
//
//     jsonFiles := map[string]string{
//         "en": "/your_json_path/en.json",
//         "ja": "/your_json_path/ja.json",
//     }
//
//   2. Initialize.
// 
//     localizedStrings.Init(jsonFiles)
//
//   3. Now you can get localized string.
//
//     3-1. When you get the localized strings in English.
//
//       strMapEn := localizedStrings.Strings("en")
//
//     3-2. When you get the localized string for the key "add" in English.
// 
//         strAddEn := localizedStrings.String("en", "add")
//
//
// MIT License
//
// Copyright (c) 2019 noknow.info
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTW//ARE. 
//////////////////////////////////////////////////////////////////////
package localizedStrings

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "runtime"
)

var (
    version = runtime.Version()
    localizedStrings = map[string]map[string]interface{}{}
)


//////////////////////////////////////////////////////////////////////
// Initialize
// @param files map[string]string: JSON files.
//        The key would be the language code.
//        The value would be the json file path.
//////////////////////////////////////////////////////////////////////
func Init(files map[string]string) {
    for lang, file := range files {
        jsonStr, err := ioutil.ReadFile(file)
        if err != nil {
            log.Fatalln(err)
        }
        mapObj := map[string]interface{}{}
        err = json.Unmarshal(jsonStr, &mapObj)
        if err != nil {
            log.Fatalln(err)
        }
        localizedStrings[lang] = mapObj
    }
}


//////////////////////////////////////////////////////////////////////
// Get localized strings.
// @param langCode string: The language code such as "en" or "ja".
//////////////////////////////////////////////////////////////////////
func Strings(langCode string) map[string]interface{} {
    targetStrings, ok := localizedStrings[langCode]
    if ok {
        return targetStrings
    } else {
        return nil
    }
}


//////////////////////////////////////////////////////////////////////
// Get the localized string.
// @param langCode string: The language code such as "en" or "ja".
// @param key string: The key of the json.
//////////////////////////////////////////////////////////////////////
func String(langCode string, key string) string {
    langStrings := Strings(langCode)
    if langStrings != nil {
        s, ok := langStrings[key].(string)
        if ok {
            return s
        } else {
            return key
        }
    } else {
        return key
    }
}
