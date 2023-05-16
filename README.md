# solar-autogtk

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>

solar-autogtk is a Go-based application that automatically sets your system's GTK theme to light or dark based on the sunrise and sunset times at your location. The application uses your system's latitude and longitude to determine the sunrise and sunset times and switches the theme accordingly, providing an adaptive environment that matches the natural lighting of your location.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

This application is built using Go. You need to have Go installed on your system. You can download and install Go from [here](https://golang.org/dl/).

You will also need the `gsettings` command available on your system. This is usually available by default on GNOME based systems.


```
sudo apt-get install gnome-tweak-tool
```

### Installing

Clone the repository:

```
git clone https://github.com/yourusername/solar-autogtk.git
```

Navigate into the project directory:

```
cd solar-autogtk
```

Compile and run the application:

```
chmod +x run.sh
./run.sh
```


## Usage <a name = "usage"></a>

You can set the light and dark themes, and your system's latitude and longitude using flags when starting the application:

```
./run.sh -light-theme Adwaita -dark-theme Adwaita-dark -latitude 40.71 -longitude -74.01 &
```

Remember about the '&' sign.

By default, the application uses "Light" and "Dark" as the light and dark theme names, and the latitude and longitude of Katowice.

The application will set your system's GTK theme to the light theme at sunrise and the dark theme at sunset. It updates the sunrise and sunset times daily at 1 am.

Please note, this application requires that the names of the light and dark themes correspond to valid GTK themes installed on your system. The latitude and longitude should be valid coordinates for accurate results.
