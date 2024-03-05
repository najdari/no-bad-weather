# No Bad Weather!

No Bad Weather is a cli application to help you dress appropriately for the weather. By simply entering a location, the app searches for the current weather conditions and, upon user confirmation or selection from multiple choices, recommends the ideal clothing for the condition.

## Public APIs

There are two public APIs that this application uses to give user the recommendation:

**1. [Geoapify](https://www.geoapify.com/)**
    To search locations and fetch associated lat and long for weather API.
    
**2. [YR](https://developer.yr.no/doc/GettingStarted/)**
    Fetch weather data of an input location.

## Running the Application

This application requires Go 1.22.0 to build and run.

### Setting Up API Key
First, you need Visit [Geoapify](https://www.geoapify.com/) and sign up to obtain your API key. 

The application and one of the test cases requires two environment variables set to function correctly:


GEOAPIFY_API_KEY=<api key generated on Geoapify>

YR_API_USER_AGENT=<set a user agent value as described [here](https://developer.yr.no/doc/TermsOfService/)>

Then you can run:
```shell
make run
```
And to run tests:
```shell
make test
```