// swift-tools-version: 6.0

import ArgumentParser
import Solutions

@main
struct AOC: AsyncParsableCommand {

    @Flag(name: .shortAndLong)
    var verbose = false

    @Argument(help: "The day for which to run the solution")
    var day: Int

    @Option(help: "The absolute folder path containing the source data")
    var folder: String?

    func run() async throws {
        let logger = ConsoleLogger(verbose: verbose)
        let solution = Solution(day: self.day, logger: logger, folder: self.folder)
        logger.log("Running day \(day) with folder \(folder ?? "nil")")

        do {
            let result = try await solution.solve()
            logger.forceLog("Result: \(result)")
        } catch (let error) {
            logger.log("Error: \(error)")
        }
    }
}
