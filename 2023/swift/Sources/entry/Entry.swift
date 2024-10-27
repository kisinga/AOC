import ArgumentParser
import solutions
import types

@main
struct Repeat: ParsableCommand {

    @Flag(name: .shortAndLong)
    var verbose = false

    @Argument(help: "The day for which to run the solution")
    var day: Int

    @Option(help: "The folder containing the source data")
    var folder: String?



    func run() throws {
        let logger = Logger(verbose: self.verbose)
        let solution = Solution(day: self.day,  logger: logger, folder: self.folder)
        logger.log("Running day \(day) with folder \(folder ?? "nil")")

        solution.run()

    }
}
