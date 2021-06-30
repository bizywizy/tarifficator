# Tarifficator

## Build

```bash
docker build -t bizywizy/tarifficator .
```

## Run

```bash
docker run --rm -it -p 8000:8000 bizywizy/tarifficator
```

## Stop

`CTRL+C`

## Usage

```bash
curl http://localhost:8000/tariffs/compare?consumption=6000
```
