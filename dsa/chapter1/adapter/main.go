package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachinAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachinAdapter)
}
