# frida-gir-comparer

Tool to compare two different frida versions by parsing frida-core.gir file.
Right now it has hardcoded kit and architecture for `macos` and `arm64`.

```bash
$ ./fgcomparer 16.4.1 16.4.8 outdir
[*] Downloading: frida-core 16.4.8 for macos arm64 to outdir/frida-core-16.4.8-macos-arm64/
[*] Downloading: frida-core 16.4.1 for macos arm64 to outdir/frida-core-16.4.1-macos-arm64/
[*] frida-core-16.4.8-macos-arm64 finished download; parsing gir file
[*] frida-core-16.4.1-macos-arm64 finished download; parsing gir file
[*] [EnumerationCount] old count: 45; new count: 47
[*] [AddedEnumerations]
	FruityDnsRecordType
	FruityDnsRecordClass
[*] [ClassCount] old count: 225; new count: 236
[*] [AddedClasses]
	FruityPairingServiceDetails
	FruityPairingStore
	FruityPairingIdentity
	FruityPairingPeer
	FruityPairingServiceMetadata
	FruityDnsPacketReader
	FruityDnsPtrRecord
	FruityDnsTxtRecord
	FruityDnsAaaaRecord
	FruityDnsSrvRecord
	FruityDnsResourceRecord
	FruityDnsResourceKey
	FruityOpackParser
	FruityMacOSCoreDeviceBackend
	FruityMacOSCoreDeviceTransport
	FruityMacOSTunnel
[*] [DeletedClasses]
	FruityMacOSFruitFinder
	FruityNullTunnelFinder
	FruityDarwinPairingBrowser
	FruityDarwinPairingServiceDetails
	FruityDarwinPairingServiceHost
```