# AVA
Ava - Another Virtual Assistant, is a modular virtual assistant built in golang. Leverages ntfy.sh and Google's Gemini; will be extended to support other LLMs in the future.

> Note:
> This project is a work in progress, as an ongoing project to learn golang. This page will update with steps to run and use Ava.

## Usage

Clone the project
```bash
git clone git@github.com:theabdullahalam/ava-go.git
```

Get a Gemini API key from [Google AI studio](https://aistudio.google.com/app/apikey?pli=1) and add the API to key to your environment:
```bash
export GEMINI_API_KEY=<YOUR_API_KEY>
```

To run Ava, in a terminal, run:
```bash
make ava
```

In another terminal, run:
```bash
make ava-cli
```
To chat with Ava.


## Limitations
- Since Ava uses the free tier of Gemini, free tier rate limits apply.

## Todo
- Run generic tasks on any nodes reliably
- Connection and disconnection of nodes (and what a node evn ultimately means in Ava's context)
- A GUI to interact with Ava
