// swift-tools-version: 6.0

import Foundation
import Types
import Utils

@frozen public enum DayXErrors: Error {
    case invalidDigitsLength
    case invalidPart
}

class DayX: Solvable {
    var filePath: String
    let logger: Logger
    let part: Int

    init(part: Int, filePath: String, logger: Logger) {
        self.filePath = filePath
        self.logger = logger
        self.part = part
    }

}

extension DayX {
    func solve() async throws -> Int {
        self.logger.log("Solving day 1 Part \(part)")
        switch part {
        case 1:
            var sum = 0

            return sum

        case 2:
            var sum = 0

            return sum
        default:
            preconditionFailure("Invalid part")
        }
    }
}
