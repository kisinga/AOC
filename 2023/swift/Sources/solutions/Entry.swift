// swift-tools-version: 6.0

import types

public class Solution {
    let day: Int
    let folder: String?
    let logger: Logger

    public init(day: Int, logger: Logger, folder: String?) {
        self.day = day
        self.folder = folder
        self.logger = logger
    }

    public func run() {
        switch day {
        case 1:
            let day1 = Day1(filePath: folder, logger: logger)
            let res = try? day1.solve()
        default:
            logger.log("Day \(day) not implemented")
        }
    }
}