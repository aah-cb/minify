# HTML minify for aah framework

Pluggable HTML Minify for aah framework.

_**Note:** This library uses `github.com/tdewolff/minify`, so it is recommended to vendor this dependencies._

## Steps to add HTML Minify

### Step 1

```bash
go get github.com/aah-cb/minify
```

### Step 2

Import into your project.
```go
import _ "github.com/aah-cb/minify"
```

### Step 3 (Optional)

If you would like to change the default config values. Add the below configuration into `render` section/block in the `aah.conf` then customize it :)

```bash
minify {
  keep {
    # To preserve all IE conditional comments such as
    # <!--[if IE 6]><![endif]--> and <![if IE 6]><![endif]>, see #https://msdn.microsoft.com/en-us/library/ms537512(v=vs.85).aspx#syntax
    # Default value is `true`
    conditional_comments = true

    # To preserve html, head and body tags
    # Default value is `true`
    document_tags = true

    # To preserve whitespace between inline tags but still collapse multiple
    # whitespace characters into one
    # Default value is `true`
    whitespace = true

    # To preserve default attribute values such as <script type="text/javascript">
    # Default value is `false`
    default_attr_vals = false

    # To preserve all end tags
    # Default value is `false`
    end_tags = false
  }
}
```

## Author

Jeevanandam M. (jeeva@myjeeva.com)

## License

`minify` released under MIT License.
