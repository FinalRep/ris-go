# RISLib – Relative Index for Streetlifting

**RISLib** ist eine Go-Bibliothek zur Berechnung des *Relative Index for Streetlifting (RIS)*. Die Bibliothek erlaubt es, den RIS-Wert anhand individueller Kraftwerte und Körpergewicht zu berechnen und die Parameter des zugrunde liegenden mathematischen Modells auf Basis realer Sportdaten zu ermitteln.

Dieses Projekt entstand im Rahmen einer Bachelorarbeit mit dem Ziel, Bewertungsverfahren im Kraftsport vergleichbar und anpassbar zu machen.

---

## 📘 Hintergrund

Der RIS ist ein relatives Bewertungssystem speziell für die Sportart **Streetlifting**. Er dient dazu, Athlet:innen unterschiedlicher Gewichtsklassen vergleichbar zu machen, indem Gesamtleistungen normalisiert werden.

### Formel

&nbsp;  
**RIS = (Total × 100) / [A + (K − A) / (1 + Q · e^(−B · (BW − v)))]**

&nbsp;

**Parameter:**

- `Total`: Gesamtleistung (z. B. Summe aus Weighted Pull-Up und Weighted Dip)
- `BW`: Körpergewicht (Bodyweight)
- `A, K, Q, B, v`: Parameter, die durch Fitting an reale Daten optimiert werden

---

## 📦 Funktionen

- Berechnung des RIS-Werts mit gegebenen Parametern
- Anpassung (Fitting) der RIS-Parameter an Leistungsdaten via nichtlinearer Optimierung
- Einlesen von CSV-Daten zur Weiterverarbeitung
- Modularer Aufbau in Go

---

## 🚀 Schnellstart

### Voraussetzungen

- Go ≥ 1.20
- [`gonum`](https://github.com/gonum/gonum) für mathematische Optimierung

### Installation

```bash
go get github.com/deinbenutzername/rislib

