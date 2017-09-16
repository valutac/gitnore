## Installation

```
go get github.com/valutac/gitnore
```
or

[Download](https://github.com/valutac/gitnore/releases/tag/0.2.0) binary file.

## Usage

Available commands:

```
update | Update      Update all gitignore configuration
list                 List available gitignore
-i | input           Gitignore input
-o | output          Gitignore output filename
```

### Basic Usage

```
$ gitnore -i go   # gitignore for golang
```

### Update Map File

```
$ gitnore update
```

### List available gitignore

```
$ gitnore list
```

Output

```
List avaiables gitignore:
actionscript, ada, agda, android, appceleratortitanium, appengine, archlinuxpackages, autotools, c, c++,
cakephp, cfwheels, chefcookbook, clojure, cmake, codeigniter, commonlisp, composer, concrete5, coq,
craftcms, cuda, d, dart, delphi, dm, drupal, eagle, elisp, elixir, elm, episerver, erlang,
expressionengine, extjs, fancy, finale, forcedotcom, fortran, fuelphp, gcov, gitbook, go, gradle, grails,
gwt, haskell, idris, igorpro, java, jboss, jekyll, joomla, julia, kicad, kohana, kotlin, labview, laravel,
leiningen, lemonstand, lilypond, lithium, lua, magento, maven, mercury, metaprogrammingsystem, nanoc, nim,
node, objective-c, ocaml, opa, opencart, oracleforms, packer, perl, phalcon, playframework, plone,
prestashop, processing, purescript, python, qooxdoo, qt, r, rails, rhodesrhomobile, ros, ruby, rust, sass,
scala, scheme, scons, scrivener, sdcc, seamgen, sketchup, smalltalk, stella, sugarcrm, swift, symfony,
symphonycms, terraform, tex, textpattern, turbogears2, typo3, umbraco, unity, unrealengine, visualstudio,
vvvv, waf, wordpress, xojo, yeoman, yii, zendframework, zephir
```

## Contributing

Pull Request is open!

## LICENSE

[MIT LICENSE](LICENSE)


