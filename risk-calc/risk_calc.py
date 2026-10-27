#!/usr/bin/env python3
"""risk-calc — консольный калькулятор кибер-рисков.

Риск = Вероятность (1–5) × Ущерб (1–5). Читает реестр угроз из CSV
(asset,threat,likelihood,impact) и печатает отсортированный реестр рисков
и матрицу 5×5.
"""
import argparse
import csv
import sys

LEVELS = [(15, "КРИТИЧЕСКИЙ"), (10, "ВЫСОКИЙ"), (5, "СРЕДНИЙ"), (0, "НИЗКИЙ")]


def level(score):
    return next(name for limit, name in LEVELS if score >= limit)


def load(path):
    with open(path, newline="", encoding="utf-8") as f:
        rows = []
        for r in csv.DictReader(f):
            l, i = int(r["likelihood"]), int(r["impact"])
            if not (1 <= l <= 5 and 1 <= i <= 5):
                sys.exit(f"ошибка: значения 1..5, строка {r}")
            rows.append({**r, "likelihood": l, "impact": i, "score": l * i})
    return sorted(rows, key=lambda r: -r["score"])


def print_register(rows):
    print(f"{'Актив':<16}{'Угроза':<28}{'В':>3}{'У':>3}{'Риск':>6}  Уровень")
    print("-" * 70)
    for r in rows:
        print(f"{r['asset']:<16}{r['threat']:<28}{r['likelihood']:>3}"
              f"{r['impact']:>3}{r['score']:>6}  {level(r['score'])}")


def print_matrix(rows):
    print("\nМатрица рисков (число угроз в ячейке), строки — вероятность:")
    print("     " + "".join(f"У={i:<3}" for i in range(1, 6)))
    for l in range(5, 0, -1):
        cells = [sum(1 for r in rows if r["likelihood"] == l and r["impact"] == i)
                 for i in range(1, 6)]
        print(f"В={l}  " + "".join(f"{c or '.':<5}" for c in cells))


def main():
    p = argparse.ArgumentParser(description="Калькулятор кибер-рисков")
    p.add_argument("csv", nargs="?", default="/app/data/threats.csv")
    p.add_argument("--min", type=int, default=0, help="показать риски >= порога")
    args = p.parse_args()

    rows = [r for r in load(args.csv) if r["score"] >= args.min]
    print_register(rows)
    print_matrix(rows)
    crit = sum(1 for r in rows if r["score"] >= 15)
    print(f"\nВсего угроз: {len(rows)}, критических: {crit}")


if __name__ == "__main__":
    main()
