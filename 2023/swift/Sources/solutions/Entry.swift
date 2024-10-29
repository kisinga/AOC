// swift-tools-version: 6.0

import Foundation
import Types
import Utils

public class Solution {
    let day: Int
    var folder: String?
    let logger: Logger

    package init(day: Int, logger: Logger, folder: String?) {
        self.day = day
        self.folder = folder
        self.logger = logger
    }

    public func solve() async throws -> Int {
        if let folder = folder {
            logger.log("Using folder \(folder)")

        } else {
            logger.log("No file path provided, using current directory")
            folder = FileManager.default.currentDirectoryPath
        }

        logger.log("Using folder \(folder)")

        switch day {
        case 1:
            let day1 = Day1(filePath: folder! + "/day1.data.txt", logger: logger)
            return try await day1.solve()
        default:
            logger.log("Day \(day) not implemented")
            return 1
        }

    }
}
