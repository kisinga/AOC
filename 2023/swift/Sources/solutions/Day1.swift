// swift-tools-version: 6.0

import Foundation
import Types
import Utils

class Day1: Solvable {
    var filePath: String
    let logger: Logger

    init(filePath: String, logger: Logger) {
        self.filePath = filePath
        self.logger = logger
    }

    func solve() async throws -> Int {
        self.logger.log("Solving day 1")
        for try await line in LineReader.lines(from: filePath) {
            logger.log(line)
        }
        return 0
    }
}
