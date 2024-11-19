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
            logger.log(.info("Using folder \(folder)"))

        } else {
            let cwd = FileManager.default.currentDirectoryPath
            logger.log(.warn("No file path provided, using current directory \(cwd)"))
            folder = cwd
        }

        switch day {
        case 1:
            let day1 = Day1(
                filePath: folder! + "/data/day1.txt", logger: logger, part: part)
            return try await day1.solve()
        case 2:
            let day2 = Day2(
                filePath: folder! + "/data/day2.txt", logger: logger, part: part)
            return try await day2.solve()
        case 3:
            let day2 = Day3(
                filePath: folder! + "/data/day3.txt", logger: logger, part: part)
            return try await day2.solve()
        default:
            logger.log(.error("Day \(day) not implemented"))
            return 1
        }

    }
}
