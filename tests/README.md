## Information

| SMBIOS Type    | Description        | Detects Enabled/Disabled?                                      |
| -------------- | ------------------ | -------------------------------------------------------------- |
| **Type 0**     | BIOS Information   | BIOS features (e.g., shadowing, boot options) via flags        |
| **Type 1**     | System Information | Mostly static data (manufacturer, serial), not feature flags   |
| **Type 4**     | Processor Info     | May indicate features like virtualization (if BIOS reports it) |
| **Type 32**    | System Boot Info   | Might contain boot status                                      |
| **Type 16/17** | Memory             | Installed or enabled memory slots                              |
| **Type 24**    | Hardware Security  | TPM, Power-on password, etc. (enabled/disabled status flags)   |
