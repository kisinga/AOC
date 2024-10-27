import ArgumentParser
import solutions
@main
struct Repeat: ParsableCommand {

    @Flag(name: .shortAndLong)
    var verbose = false

    @Option(name: [.short, .long], help: "The day for which to run the solution")
    var day: Int

    @Argument(help: "The folder containing the source data")
    var folder: String?

    let logger = Logger()

    mutating func run() throws {
        if verbose {
            print("Running with verbose output...")
        }
        print("Running day \(day) with folder \(folder ?? "nil")")



    }
}
