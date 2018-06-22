To create your project, start by creating a *.cutedoc.yml* file in a directory of your choice, usually
the in the root of your project. This is the general structure of a manifest file:
```
meta:
  title: Cutedoc
  author: NickTheDev
  description: Cutedoc's documentation.
  branding:
    icon: docs/static/img/icon.png
    logo: docs/static/img/logo.png
articles:
  Introduction:
    Intro: docs/sources/about.md
```
### Meta

To describe basic characteristics of the project, you can define a meta section in the manifest
file. In this section you may define meta such as the title, author, and description and the
project branding.
```
# Project metadata.
meta:
  # Html header metadata title. Default is blank.
  title: Cutedoc

  # Html header metadata author. Default is blank.
  author: NickTheDev

  # Html header metadata description. Default is blank.
  description: Cool documentation.

  # Project branding.
  branding:
    # Html header favicon, preferably 16 x 16. Default is null.
    icon: path/to/icon.png

    # Logo of the project, preferably 200 x 100. Default is null.
    logo: path/to/logo.png
```
### Build

To configure the build of your project, you can define a build section in the manifest file. In this section you may
specify the output directory of the build and whether to minify the production assets or not.
```
# Build config.
build:
  # Output directory of the documentation. Default is 'docs'.
  dir: docs

  # Whether or not to minify the build assets (css, js). Default is true.
  minify: true
```
### Articles

To link markdown files in your project, you can define an articles section in the manifest file.
In this section you may define article groups, and then place the links in them.
```
# Project articles.
articles:
  # Article group name.
  Introduction:
    # Path to article markdown file, may use github flavored markdown.
    Intro: docs/sources/about.md
```
### Theme
To configure the theme of your project, you can define a theme section in the manifest file.
You may specify the template name, the only one currently being the minimal theme (The
documentation you are reading right now uses the minimal theme). You can also
define the color scheme of the template, namely the primary, secondary, text, nav, and
background colors.
```
# Project theme.
theme:
  # Theme template, currently the only supported template is 'minimal'. Default is 'minimal'.
  template: minimal

  # Template color scheme.
  colors:
    # Primary color used for headers and navigation group text colors. Default is 3e4669.
    primary: 3e4669

    # Secondary color used an accent in the navigation and the body headers. Default is c3c6de.
    secondary: c3c6de

    # Body text color used in the documentation. Default is 3e4669.
    text: 3e4669

    # Navigation background color. Default is fff.
    nav: fff

    # Background color. Default is f6f7fb.
    background: f6f7fb
```