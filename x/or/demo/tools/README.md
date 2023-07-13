# Tools

- cannon - mips vm, execute step, generate state and witness.
- challenge - challenge emulator, output id, steps, and the flag if challenge is searching.
- miniapp - verification program, verify if sequecner's state is correct.
- preimage - preimage oracle, serve data for miniapp.

# Preimage/Data
 - data/height-100.json - This is correct data.
 - data/height-100-defender.json - This is malicious data. The demo assume that this data is proposed by malicious sequencer.
 - data/height-100-challenge.json - This data is used in fraud proof. Challenger prove sequencer's fraud using this data.
 