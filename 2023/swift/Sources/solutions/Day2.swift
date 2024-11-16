// swift-tools-version: 6.0

import Foundation
import Types
import Utils

@frozen public enum Day2Errors: Error {
    case invalidDigitsLength
    case invalidPart
}

struct Day2: Solvable {
    var filePath: String
    let logger: Logger
    let part: Int
}

struct Game {
    let red: Int
    let green: Int
    let blue: Int
}

extension Day2 {
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
