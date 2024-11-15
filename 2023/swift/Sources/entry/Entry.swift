// swift-tools-version: 6.0

import ArgumentParser
import Solutions

@main
struct AOC: AsyncParsableCommand {

    @Flag(name: .shortAndLong)
    var verbose = false

    @Option(help: "The day for which to run the solution")
    var day: Int

    @Option(help: "The absolute folder path containing the source data")
    var folder: String?

    @Option(help: "part 1 or 2 ONLY")
    var part: Int

    func run() async throws {
        precondition((part > 0 && part < 3), "Part must be 1 or 2")
        let logger = ConsoleLogger(verbose: verbose)
        let solution = Solution(for: day, part: part, logger: logger, folder: self.folder)
        logger.log("Running day \(day) with folder \(folder ?? "nil")")

        do {
            let result = try await solution.solve()
            logger.forceLog("Result: \(result)")
        } catch (let error) {
            logger.log("Error: \(error)")
        }
    }
}
