# Hola Mundo en Go

Este documento te guía para crear y ejecutar un programa básico de "Hola Mundo" en Go.

## Requisitos 🟠

- Tener instalado [Go](https://golang.org/dl/).
- Un editor de texto (VS Code, Sublime, etc).

## Pasos para crear el programa

1. **Crea una carpeta para tu proyecto:**
	```sh
	mkdir hola-mundo
	cd hola-mundo
	```

2. **Crea el archivo principal:**
	```sh
	code main.go
	```
	O usa tu editor favorito para crear `main.go`.

3. **Escribe el código:**
	```go
	package main

	import "fmt"

	func main() {
		 fmt.Println("Hola Mundo")
	}
	```

## Ejecutar el programa

1. Abre una terminal en la carpeta del proyecto.
2. Ejecuta el siguiente comando:
	```sh
	go run main.go
	```

Deberías ver en pantalla:
```
Hola Mundo
```

## Compilar el programa (opcional)

Si quieres crear un ejecutable:
```sh
go build main.go
```
Esto generará un archivo `main.exe` (Windows) o `main` (Linux/Mac).

## Recursos útiles

- [Documentación oficial de Go](https://golang.org/doc/)
- [Tour de Go](https://tour.golang.org/)



## Tablas

|Datos |🟢 wifi|🟢 bluethod|🟢 nfc|🟢 bateria|🔴 test|🔴 test|🟣 test
|---|:---:|:---:|:---:|:---:|:---:|:---:|---|
|1. T Android|❌|❌|❌|❌|❌|❌|test
|2. C Iphone✔️|✔️|✔️|✔️|✔️|✔️|test



### 🟢 Ticket
(un escaneo) Tipo de ticket comun para un evento

### 🟢 Mesa
(un escaneo) Tipo de ticket con un lugar especifico de mesa

### 🟢 Cortesía
 (un escaneo) Tipo de ticket con mension especifica del tipo de entrada cortesía. No se incluyo sector debido a que este se asocia con uno o mas sectores a los que puede entrar

### 🟢 Abono
(un escaneo al dia) Tipo de ticket recurrente que sirve para varios dias. Incluye en vez del campo Fecha Validez un rango de fechas

### 🔴 Butaca
(un escaneo) Tipo de ticket que es para un sitio especifico

### 🔴 Físico
(un escaneo) Tipo de ticket para venta física

#### 🟣 Longitud
Tamaño minimo y maximo de los caracteres que se puede incluir en estos campos

# Test

# Headings
To create a heading, add number signs (#) in front of a word or phrase. The number of number signs you use should correspond to the heading level. For example, to create a heading level three (<h3>), use three number signs (e.g., ### My Header).

# Heading level 1
Heading level 1
## Heading level 2
Heading level 2
### Heading level 3
Heading level 3
#### Heading level 4
Heading level 4
##### Heading level 5
Heading level 5
###### Heading level 6

# Alternate Syntax
Alternatively, on the line below the text, add any number of == characters for heading level 1 or -- characters for heading level 2.

Heading level 1
===============

Heading level 2
---------------

# Paragraphs
To create paragraphs, use a blank line to separate one or more lines of text.

I really like using Markdown.

I think I'll use it to format all of my documents from now on.

# Line Breaks
To create a line break or new line (<br>), end a line with two or more spaces, and then type return.

This is the first line.  
And this is the second line.

# Emphasis
You can add emphasis by making text bold or italic.

## Bold
To bold text, add two asterisks or underscores before and after a word or phrase. To bold the middle of a word for emphasis, add two asterisks without spaces around the letters.

I just love **bold text**.
I just love __bold text__.
Love**is**bold

## Italic
To italicize text, add one asterisk or underscore before and after a word or phrase. To italicize the middle of a word for emphasis, add one asterisk without spaces around the letters.

Italicized text is the *cat's meow*.
Italicized text is the _cat's meow_.
A*cat*meow

## Bold and Italic
To emphasize text with bold and italics at the same time, add three asterisks or underscores before and after a word or phrase. To bold and italicize the middle of a word for emphasis, add three asterisks without spaces around the letters.

This text is ***really important***.
This text is ___really important___.
This text is __*really important*__.
This text is **_really important_**.
This is really***very***important text.

# Blockquotes
To create a blockquote, add a > in front of a paragraph.

> Dorothy followed her through many of the beautiful rooms in her castle.

# Blockquotes with Multiple Paragraphs

> Dorothy followed her through many of the beautiful rooms in her castle.
>
> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.

# Nested Blockquotes
Blockquotes can be nested. Add a >> in front of the paragraph you want to nest.

> Dorothy followed her through many of the beautiful rooms in her castle.
>
>> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.

# Blockquotes with Other Elements
Blockquotes can contain other Markdown formatted elements. Not all elements can be used — you’ll need to experiment to see which ones work.

> #### The quarterly results look great!
>
> - Revenue was off the chart.
> - Profits were higher than ever.
>
>  *Everything* is going according to **plan**.


# Lists
You can organize items into ordered and unordered lists.

## Ordered Lists
To create an ordered list, add line items with numbers followed by periods. The numbers don’t have to be in numerical order, but the list should start with the number one.

1. First item
2. Second item
3. Third item
4. Fourth item

1. First item
1. Second item
1. Third item
1. Fourth item

1. First item
8. Second item
3. Third item
5. Fourth item

1. First item
2. Second item
3. Third item
    1. Indented item
    2. Indented item
4. Fourth item


# Unordered Lists
To create an unordered list, add dashes (-), asterisks (*), or plus signs (+) in front of line items. Indent one or more items to create a nested list.

- First item
- Second item
- Third item
- Fourth item

* First item
* Second item
* Third item
* Fourth item

+ First item
+ Second item
+ Third item
+ Fourth item

- First item
- Second item
- Third item
    - Indented item
    - Indented item
- Fourth item


# Starting Unordered List Items With Numbers
If you need to start an unordered list item with a number followed by a period, you can use a backslash (\) to escape the period.

- 1968\. A great year!
- I think 1969 was second best.


# Adding Elements in Lists
To add another element in a list while preserving the continuity of the list, indent the element four spaces or one tab, as shown in the following examples.

## Paragraphs

* This is the first list item.
* Here's the second list item.

    I need to add another paragraph below the second list item.

* And here's the third list item.

## Blockquotes

* This is the first list item.
* Here's the second list item.

    > A blockquote would look great below the second list item.

* And here's the third list item.

## Code Blocks

1. Open the file.
2. Find the following code block on line 21:

        <html>
          <head>
            <title>Test</title>
          </head>

3. Update the title to match the name of your website.

## Images

1. Open the file containing the Linux mascot.
2. Marvel at its beauty.

    ![Tux, the Linux mascot](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRczg28ObvWILeiI7fUW5U1o1nkajUeWMsoGw&s)
	

3. Close the file.

## Lists
You can nest an unordered list in an ordered list, or vice versa.

1. First item
2. Second item
3. Third item
    - Indented item
    - Indented item
4. Fourth item

## Code
To denote a word or phrase as code, enclose it in backticks (`).

At the command prompt, type `nano`.

## Escaping Backticks
If the word or phrase you want to denote as code includes one or more backticks, you can escape it by enclosing the word or phrase in double backticks (``).

``Use `code` in your Markdown file.``

## Code Blocks
To create code blocks, indent every line of the block by at least four spaces or one tab.

   <html>
      <head>
      </head>
    </html>

# Horizontal Rules
To create a horizontal rule, use three or more asterisks (***), dashes (---), or underscores (___) on a line by themselves.


***

---

_________________


# Links
To create a link, enclose the link text in brackets (e.g., [Duck Duck Go]) and then follow it immediately with the URL in parentheses (e.g., (https://duckduckgo.com)).

My favorite search engine is [Duck Duck Go](https://duckduckgo.com).


# Adding Titles
You can optionally add a title for a link. This will appear as a tooltip when the user hovers over the link. To add a title, enclose it in quotation marks after the URL.

My favorite search engine is [Duck Duck Go](https://duckduckgo.com "The best search engine for privacy").

# URLs and Email Addresses
To quickly turn a URL or email address into a link, enclose it in angle brackets.

<https://www.markdownguide.org>
<fake@example.com>

# Formatting Links
To emphasize links, add asterisks before and after the brackets and parentheses. To denote links as code, add backticks in the brackets.

I love supporting the **[EFF](https://eff.org)**.
This is the *[Markdown Guide](https://www.markdownguide.org)*.
See the section on [`code`](#code).

## An Example Putting the Parts Together
Say you add a URL as a standard URL link to a paragraph and it looks like this in Markdown:

In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, filled with the ends
of worms and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to
eat: it was a [hobbit-hole](https://en.wikipedia.org/wiki/Hobbit#Lifestyle "Hobbit lifestyles"), and that means comfort.


In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, filled with the ends
of worms and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to
eat: it was a [hobbit-hole][1], and that means comfort.

[1]: <https://en.wikipedia.org/wiki/Hobbit#Lifestyle> "Hobbit lifestyles"


Images
To add an image, add an exclamation mark (!), followed by alt text in brackets, and the path or URL to the image asset in parentheses. You can optionally add a title in quotation marks after the path or URL.

![The San Juan Mountains are beautiful!](/assets/images/shiprock.jpg "San Juan Mountains")

# Linking Images
To add a link to an image, enclose the Markdown for the image in brackets, and then add the link in parentheses.

[![An old rock in the desert](/assets/images/shiprock.jpg "Shiprock, New Mexico by Beau Rogers")](https://www.flickr.com/photos/beaurogers/31833779864/in/photolist-Qv3rFw-34mt9F-a9Cmfy-5Ha3Zi-9msKdv-o3hgjr-hWpUte-4WMsJ1-KUQ8N-deshUb-vssBD-6CQci6-8AFCiD-zsJWT-nNfsgB-dPDwZJ-bn9JGn-5HtSXY-6CUhAL-a4UTXB-ugPum-KUPSo-fBLNm-6CUmpy-4WMsc9-8a7D3T-83KJev-6CQ2bK-nNusHJ-a78rQH-nw3NvT-7aq2qf-8wwBso-3nNceh-ugSKP-4mh4kh-bbeeqH-a7biME-q3PtTf-brFpgb-cg38zw-bXMZc-nJPELD-f58Lmo-bXMYG-bz8AAi-bxNtNT-bXMYi-bXMY6-bXMYv)


# Escaping Characters
To display a literal character that would otherwise be used to format text in a Markdown document, add a backslash (\) in front of the character.

\* Without the backslash, this would be a bullet in an unordered list.

# HTML
Many Markdown applications allow you to use HTML tags in Markdown-formatted text. This is helpful if you prefer certain HTML tags to Markdown syntax. For example, some people find it easier to use HTML tags for images. Using HTML is also helpful when you need to change the attributes of an element, like specifying the color of text or changing the width of an image.

This **word** is bold. This <em>word</em> is italic.