# Building Cross-Platform GUI Applications with Fyne

<a href="https://www.packtpub.com/product/building-cross-platform-gui-applications-with-fyne/9781800563162?utm_source=github&utm_medium=repository&utm_campaign=9781800563162"><img src="https://static.packt-cdn.com/products/9781800563162/cover/smaller" alt="Building Cross-Platform GUI Applications with Fyne" height="256px" align="right"></a>

This is the code repository for [Building Cross-Platform GUI Applications with Fyne](https://www.packtpub.com/product/building-cross-platform-gui-applications-with-fyne/9781800563162?utm_source=github&utm_medium=repository&utm_campaign=9781800563162), published by Packt.

**Create beautiful, platform-agnostic graphical applications using Fyne and the Go programming language**

## What is this book about?
The history of graphical application development is long and complicated, with various development challenges that persist to this day. The mix of technologies involved and the need to use different programming languages lead to a very steep learning curve for developers looking to build applications across multiple platforms. 

This book covers the following exciting features:
Become well-versed with the history of GUI development and how Fyne and the Golang programming language make it easier
Explore how the Fyne toolkit is architected and the various modules are provided
Discover how Fyne apps can be tested and constructed using best practices
Construct five complete applications and deploy them to your devices
Customize the design of your apps by extending widgets and themes
Understand the separation and presentation of data and how to test and build applications that present dynamic data

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1800563167) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```
package main
import (
    "log"
    "os/exec"
)
```

**Following is what you need for this book:**
This Fyne-Golang GUI book is for developers from any background who are looking to build cross-platform applications with a modern toolkit. It will also be useful for Go developers who are looking to explore graphical apps and GUI developers looking for a new toolkit for cross-platform development. Basic knowledge of Graphical User Interface (GUI) development is assumed (although a brief history is also included in the book). The book also features a short introduction to the Go language as a quick refresher.	

With the following software and hardware list you can run all code files present in the book (Chapter 1-10).
### Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 1 | Go 1.12 (or newer) | Windows, Mac OS X, and Linux (Any) |
| 2 | Fyne 2.0 | iOS, iPad or Android |

We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it](https://static.packt-cdn.com/downloads/9781800563162_ColorImages.pdf).

### Related products
* Mastering Go - Second Edition [[Packt]](https://www.packtpub.com/product/mastering-go-second-edition/9781838559335?utm_source=github&utm_medium=repository&utm_campaign=9781838559335) [[Amazon]](https://www.amazon.com/dp/1838559337)

* Hands-On High Performance with Go [[Packt]](https://www.packtpub.com/product/hands-on-high-performance-with-go/9781789805789?utm_source=github&utm_medium=repository&utm_campaign=9781789805789) [[Amazon]](https://www.amazon.com/dp/B08576P94D)


## Corrections

### V2 imports

In the first edition of this book the import path for fyne is listed as `fyne.io/fyne`. However since the January release of v2.0.0 the imports should be `fyne.io/fyne/v2`.

For example, on page 49 the line:

```go
import "fyne.io/fyne/app"
```

should read

```go
import "fyne.io/fyne/v2/app"
```

In addition to the code changes commands that install Fyne tools are impacted, so to install the `fyne` tool the correct command on page 52 would be:

```
go get fyne.io/fyne/v2/cmd/fyne_demo
```

With these changes applied, or by using the code directly in this repository, the examples will all operate as described in the book content.

### Chapter 5

Page 121 example of pop-up menu code should read:

```go
menu := fyne.NewMenu("",
        fyne.NewMenuItem("Item 1", func() {}),
        fyne.NewMenuItem("Item 2", func() {}),
    )
    pos := fyne.NewPos(20, 20)
    widget.ShowPopUpMenuAtPosition(menu, w.Canvas(), pos)
```
### Chapter 7

References to `ThemeColourName` should read `ThemeColorName`


### Appendix C: Cross-Compiling

On pages 282 and 283 where macOS cross-compiling commands are described they assume an Intel based Apple device.
When working on an M1 or M2 processor you will beed to add a `GOARCH=amd64` parameter alongside `GOOS` to specify an intel based 64bit application build.


## Get to Know the Author
**Andrew Williams** graduated from the University of Edinburgh in 2003 with a bachelor’s degree, with honors, in computer science. After university, he went to work as a software engineer and has gained over 15 years of commercial software development experience across a variety of programming languages, including Java, C, Objective-C, and Go. Andrew has spent many years working as a CTO with many early-stage and growing software start-ups. He has been a core developer in large open source projects, including Enlightenment, EFL, and Maven, as well as involved in maintaining various community websites and tutorials. Andrew’s passion for building tools and services that make software development simpler led him to start authoring books on the subject.



