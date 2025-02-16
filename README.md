This is the centralized error handling in golang

- We got a middleware which catch each panic and recover them.
- The panic can get an error which the middleware will log it using the logger into a file with the date and line of the error which is usefull in production.
- The panic can be a custom message which got a MessageKey (Which also can be handled with text keys instead of numuric keys), this way the middleware will bring the specific error message from the template and return it.
  
