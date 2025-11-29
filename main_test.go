package main

import (
 "testing"
)

func TestWorkerSort(t *testing.T) {
 c := &Company{}

 // Добавляем работников
 c.AddWorkerInfo("Михаил", "директор", 3000, 4)         // 12000
 c.AddWorkerInfo("Андрей", "мастер", 1800, 6)           // 10800
 c.AddWorkerInfo("Игорь", "зам. директора", 1800, 3)    // 5400
 c.AddWorkerInfo("Николай", "начальник цеха", 720, 4)   // 2880
 c.AddWorkerInfo("Виктор", "рабочий", 720, 4)           // 2880

 expected := []string{
  "Михаил — 12000 — директор",
  "Андрей — 10800 — мастер",
  "Игорь — 5400 — зам. директора",
  "Николай — 2880 — начальник цеха",
  "Виктор — 2880 — рабочий",
 }

 got, err := c.SortWorkers()
 if err != nil {
  t.Fatalf("SortWorkers вернул ошибку: %v", err)
 }

 if len(got) != len(expected) {
  t.Fatalf("Размер результата неверный: ожидалось %d, получили %d", len(expected), len(got))
 }

 for i := range expected {
  if got[i] != expected[i] {
   t.Fatalf("SortWorkers неправильно отсортировал.\nОжидалось: %v\nПолучено: %v", expected, got)
  }
 }
}