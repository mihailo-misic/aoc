#!/bin/bash
set -euo pipefail

ROOT_DIR="$(dirname $(dirname "$(realpath "$0")"))"

CUR_YEAR="$(date +%Y)"
read -p "Enter year [$CUR_YEAR]: " YEAR
YEAR=${YEAR:-$CUR_YEAR}
YEAR_DIR="$ROOT_DIR/$YEAR"
mkdir -p $YEAR_DIR

LATEST_DAY=$(ls -d "$YEAR_DIR/*/" 2>/dev/null | grep -Eo '[0-9]+' | sort -n | tail -n 1) || LATEST_DAY="00"
NEXT_DAY=$((10#$LATEST_DAY + 1))
read -p "Enter day  [$NEXT_DAY]: " DAY
DAY=${DAY:-$NEXT_DAY}
DAY=$(printf "%02d" $DAY)
DAY_DIR="$YEAR_DIR/$DAY"
mkdir -p $DAY_DIR

cp -r --update=none "$ROOT_DIR/seed"/* "$DAY_DIR"


# Get input from aoc
INPUT_URL="https://adventofcode.com/$YEAR/day/$((10#$DAY))/input"
echo -e "\n\033[0;32m[GET] $INPUT_URL\033[0m"
curl --cookie "session=$AOC_SESSION" $INPUT_URL -o "$DAY_DIR/input.txt"
echo -e "\nInput:\n..."
tail -n 10 "$DAY_DIR/input.txt"

