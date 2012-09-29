gotamer/cfg
###########

### gotamer/cfg is a json configuration package
You define your config items in your application as a struct.  
If the config file doesn't exist, gotamer/cfg will create a json
template file for you, according to the struct definition in your application.

		type Sqlite struct {
			Path string
		}
		cfg.APPL     = "CafeMaker"
		cfg.FileName = "database"
		config := cfg.Get(Sqlite)


[Website](http://www.robotamer.com/html/GoTamer/cfg.html "gotamer/cfg Documentation")

________________________________________________________

The MIT License (MIT)

Copyright © 2012 GoTamer <http://www.robotamer.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
