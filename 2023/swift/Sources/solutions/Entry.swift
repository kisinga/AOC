// swift-tools-version: 6.0

import Foundation
import Types
import Utils

public class Solution {
    let day: Int
    let part: Int
    var folder: String?
    let logger: Logger

    package init(for day: Int, part: Int, logger: Logger, folder: String?) {
        self.day = day
        self.folder = folder
        self.logger = logger
        self.part = part
    }

    public func solve() async throws -> Int {
        if let folder {
            logger.log("Using folder \(folder)")

        } else {
            let cwd = FileManager.default.currentDirectoryPath
            logger.log("No file path provided, using current directory \(cwd)")
            folder = cwd
        }

        switch day {
        case 1:
            let day1 = Day1(
                filePath: folder! + "/challenges/day1.data.txt", logger: logger, part: part)
            return try await day1.solve()
        default:
            logger.log("Day \(day) not implemented")
            return 1
        }

    }
}
