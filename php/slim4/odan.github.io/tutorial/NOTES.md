#Slim 4 Tutorial:

I really like some patterns included in this tutorial, although I'm not sure 
about 'single action' services vs multi action models. The 
'domain' approach in general though feels much nicer than C.R.U.D. 
models and I like the user of 'repositories' to access different datastores.

Using immediately executing functions when requiring files to load settings etc. 
is also interesting. This feels like a nice way to load data and analogous to the 
use of modules in JS.

I'm not entirely comfortable with callabales for classes in PHP
but seem to have got my head aroung the `__invoke()` magic method. 
This is an area I need to spend me time learning about though.

Autowiring through the PHP-DI library again feels magic... I 
assume this is wrapping PHP/Composers autoload function somehow.

Despite starting the tutorial to get a better understanding of 
middleware in Slim4 it didn't really cover the use of middleware, 
specifically "on the way in" as opposed to out. Still I feel I 
have a better appreciation for how to structure a web application 
than I did before.
