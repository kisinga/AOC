// swift-tools-version: 6.0

import Foundation
import Types
import Utils

@frozen public enum DayXErrors: Error {
    case invalidDigitsLength
    case invalidPart
}

struct DayX: Solvable {
    var filePath: String
    let logger: Logger
    let part: Int
}

extension DayX {
    func solve() async throws -> Int {
        switch part {
        case 1:
            var sum = 0
            for try await line in FileReader.lines(from: filePath) {

            }
            return sum

        case 2:
            var sum = 0

            return sum
        default:
            preconditionFailure("Invalid part")
        }
    }
}
