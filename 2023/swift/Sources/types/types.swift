class Logger {
    var verbose: Bool = false
    func log(_ message: String) {
        if verbose {
            print(message)
        }
    }
}
