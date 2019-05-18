# Visualisierung für Sortieralgorithmen

## Was macht dieses Tool?
Mit Hilfe dieses Tools können Sortieralgorithmen auf Konsolenebene visualisiert werden.
Zur Visualisierung wird hierfür tcell[github.com/gdamore/tcell] genutzt.

## Getting started

Zunächst das Github Repository klonen:

```
git clone github.com/ob-algdatii-ss19/leistungsnachweis-lallinger-stortz
```

Anschließend bauen des Tools mit:

```
go build
```

Danach kann das Tool wie im nächsten Abschnitt beschrieben ausgeführt werden.

## Usage
Es kann ein Sortieralgorithmus und ein stepping mode (manuell oder automatisch) gewählt werden.
Während automatisches stepping an ist ist es möglich mit Leertaste in den manuellen Modus zu wechseln.
Im manuellen Modus kann ein Schritt durch drücken der Leertaste ausgeführt werden oder durch drücken von 'c' zurück in den automatischen Modus gewechselt werden.
Um das Programm zu beenden kann jederzeit strg+c gedrückt werden.
Ein Ändern der Fenstergröße während der Ausführung ist nicht möglich. Bitte setzen Sie die gewünschte Fenstergröße bevor Sie das Programm ausführen.

```
leistungsnachweis-lalllinger-stortz [FLAGS] SORTER
FLAGS
Usage of C:\Users\Patrick\Desktop\SkyDrive\Studium\6_SS_19\Algo_Dat_2\Praktikum\leistungsnachweis-lallinger-stortz\leistungsnachweis-lallinger-stortz.exe:
  -sleep int
        Time to wait between steps while automatic stepping in miliseconds (default 100)
  -step
        Activate manual stepping

SORTER
  bubble
        Sort values using bubblesort
  insertion
        Sort values using insertionsort
  selection
        Sort values using selection
  bogo
        Sort values using bogosort. Caution this probably won't terminate by itself!

Example:
  leistungsnachweis-lalllinger-stortz --sleep=50 --step=false selection
```
