package main

import (
	"archive/zip"
	"fmt"
	"os"
	"strings"
)

// Framework define el tipo para los nombres de los frameworks
type Framework string

// Constantes para los nombres de los diferentes frameworks
//TODO-VALIDATE FRAMEWORKS
const (
	Flutter       Framework = "Flutter"
	ReactNative   Framework = "React Native"
	Cordova       Framework = "Cordova"
	Xamarin       Framework = "Xamarin"
	Ionic         Framework = "Ionic"
	Native        Framework = "Native (Java/Kotlin)"
	Swift         Framework = "Swift (iOS)"
	ObjectiveC    Framework = "Objective-C (iOS)"
	Unity         Framework = "Unity"
	PhoneGap      Framework = "PhoneGap"
	Angular       Framework = "Angular"
	Vue           Framework = "Vue.js"
	Svelte        Framework = "Svelte"
)

// Technology asocia un framework con directorios característicos
type Technology struct {
	Framework   Framework
	Directories []string
}

// techList contiene las tecnologías y los directorios específicos que generalmente se encuentran en sus paquetes.
var techList = []Technology{
	{Framework: Flutter, Directories: []string{"libflutter.so"}},
	{Framework: ReactNative, Directories: []string{"libreactnativejni.so", "assets/index.android.bundle"}},
	{Framework: Cordova, Directories: []string{"assets/www/index.html", "assets/www/cordova.js", "assets/www/cordova_plugins.js"}},
	{Framework: Xamarin, Directories: []string{"/assemblies/Sikur.Monodroid.dll", "/assemblies/Sikur.dll", "/assemblies/Xamarin.Mobile.dll", "/assemblies/mscorlib.dll", "libmonodroid.so", "libmonosgen-2.0.so"}},
	{Framework: Ionic, Directories: []string{"assets/www/build/main.js", "assets/www/index.html", "assets/www/cordova.js"}},
}

// main es el punto de entrada del programa.
func main() {
	// Obtiene la ruta del archivo .apk o .ipa desde los argumentos de la línea de comandos.
	appPath := getAppPath()

	// Intenta abrir el archivo como un archivo zip.
	zipFile, err := zip.OpenReader(appPath)
	if err != nil {
		fmt.Printf("Error opening zip file: %s\n", err)
		return
	}
	defer zipFile.Close()

	// Busca en los archivos del zip signos de las tecnologías definidas.
	for _, file := range zipFile.File {
		for _, tech := range techList {
			for _, directory := range tech.Directories {
				if strings.Contains(file.Name, directory) {
					fmt.Printf("App was written in %s\n", tech.Framework)
					return // Finaliza tan pronto como encuentre una coincidencia
				}
			}
		}
	}
	// Si no se encuentra ninguna tecnología, se asume que es nativa.
	fmt.Printf("App was written in %s\n", Native)
}

// getAppPath recupera la ruta del archivo de la aplicación desde los argumentos de la línea de comandos.
func getAppPath() string {
	if len(os.Args) > 1 {
		return os.Args[1] // Retorna la ruta pasada como argumento.
	}
	// Si no se proporciona la ruta, imprime un mensaje de error y termina el programa.
	fmt.Println("Please provide the full path to the .ipa or .apk file as an argument." +
		"\nEg: go run main.go /path/to/app.apk")
	os.Exit(1)
	return ""
}
