// swift-tools-version: 6.0

import Foundation
import Types
import Utils

class Day1: Solvable {
    var filePath: String?
    let logger: Logger

    init(filePath: String?, logger: Logger) {
        self.filePath = filePath
        self.logger = logger
    }

    func solve() throws -> Int {
        self.logger.log("Solving day 1")
        if let filePath = self.filePath {
            logger.log("Using file \(filePath)")
        } else {
            logger.log("No file path provided, using current directory")
            filePath = FileManager.default.currentDirectoryPath
        }
        return 0
    }
}
