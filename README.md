# tokinfo

Herramienta CLI para mejorar prompts usando Gemini AI.

## Resumen

`tokinfo` es una herramienta de línea de comandos escrita en Go diseñada para mejorar prompts de usuario. Utiliza el modelo Gemini AI para aplicar directrices de ingeniería de prompts definidas en un archivo de configuración JSON, generando prompts más efectivos y refinados.

## Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/JuanJml01/Improved-prompt
   ```
2. Navega al directorio del proyecto:
   ```bash
   cd tokinfo
   ```
3. Compila el ejecutable:
   ```bash
   go build
   ```
   Esto creará un archivo ejecutable llamado `tokinfo` en el directorio actual.


## Uso

Para usar la herramienta después de compilar:
```bash
./tokinfo "Tu prompt inicial aquí"
```
Para usar la herramienta después de instalar con `go install`:
```bash
tokinfo "Tu prompt inicial aquí"
```
La herramienta procesará tu prompt utilizando las directrices del archivo `guidelines.json` (o el archivo de configuración especificado) y la API de Gemini AI, imprimiendo el prompt mejorado en la salida estándar.

## Descripción

`tokinfo` es una herramienta CLI en Go que mejora prompts usando Gemini AI y directrices JSON. Permite aplicar técnicas de ingeniería de prompts consistentemente.

## Limitaciones

El proyecto está en desarrollo activo y puede contener errores o limitaciones. Se proporciona "tal cual" sin garantías.