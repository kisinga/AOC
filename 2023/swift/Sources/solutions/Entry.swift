// swift-tools-version: 6.0

import Types
import Utils

public class Solution {
    let day: Int
    let folder: String?
    let logger: Logger

    package init(day: Int, logger: Logger, folder: String?) {
        self.day = day
        self.folder = folder
        self.logger = logger
    }

    public func solve() {
        switch day {
        case 1:
            let day1 = Day1(filePath: folder, logger: logger)
            let res = try? day1.solve()
        default:
            logger.log("Day \(day) not implemented")
        }
    }
}
