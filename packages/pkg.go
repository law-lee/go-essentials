package packages

// You can also access local packages inside of the current folder by 
// prefacing packages with ./. In a project with a structure like this:
/*
import (
    "./src/package1"
    "./src/package2"
)
*/
/*
Since package-names can collide in different libraries you may want to alias one package to a new name. You can do this by prefixing your import-statement with the name you want to use.

import (
    "fmt" //fmt from the standardlibrary
    tfmt "some/thirdparty/fmt" //fmt from some other library
)
This allows you to access the former fmt package using fmt.* and the latter fmt package using tfmt.*.


It is considered a bad style but you can import the package into the global namespace to avoid using package prefix:

import (
    . "fmt"
)
With the above Printf is the same as fmt.Printf.


Go compiler reports an error if you import a package but don't use it. To work-around this, you can set the alias to the underscore:

import (
    _ "fmt"
)
This might be necessary if you don't use the package directly but need the side-effects of running package's init function.
*/