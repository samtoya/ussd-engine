# USSD Builder Gateway
This project is a gateway into the world of USSD to handle any USSD providers.

To simply put, it's core design inspiration was taken from years of development of the USSD flow
and majority from the Django project "django-airflow"

### Screen Types
1. initial_screen
2. input_screen
3. menu_screen
4. quit_screen


## Initial Screen
Type: initial_screen
Description: All and every application must have and start with this screen. This initializes the application.
Example:
```yaml
initial_screen: welcome_screen
type: initial_screen
text: This is the very first screen displayed to the user
next_screen: menu_screen
```



#### Development

```go
type UssdScreens struct {
	Screens map[string]*Screen
}

yaml = map[string]*UssdScreen
```

Where the key of type string is the name of the screen to navigate and the screen holds the core functionality of that screen. Within the entire map,
there must be only one screen of type initial_screen.