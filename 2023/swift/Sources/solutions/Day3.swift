// swift-tools-version: 6.0

import Foundation
import Types
import Utils

@frozen public enum Day3Errors: Error {
    case invalidDigitsLength
    case invalidPart
}

struct Day3: Solvable {
    var filePath: String
    let logger: Logger
    let part: Int
}

struct NumberCluser {
    let digits: String
    let startPosition: Int
    let length: Int
}

class LineProcessor {
    private var values: [String]
    init() {
        self.values = Array(repeating: "", count: 3)
    }

    func process(_ value: String) -> Int? {
        if current.isEmpty {
            values[1] = value
        } else if previous.isEmpty {
            values.removeFirst()
            values[1] = value
        } else if next.isEmpty {
            values[2] = value
        } else {
            values.removeFirst()
            values[2] = value
        }
        return 0
    }

    private var current: String {
        return values[1]
    }

    private var previous: String {
        return values[0]
    }

    private var next: String {
        return values[2]
    }
}

func findDigitClusters(in text: String) -> [NumberCluser] {
    var results: [NumberCluser] = []

    for match in text.matches(of: /\d+/) {
        let digits = String(match.0)
        let startPosition = text.distance(from: text.startIndex, to: match.startIndex)
        let length = digits.count

        results.append(
            NumberCluser(
                digits: digits,
                startPosition: startPosition,
                length: length
            ))
    }

    return results
}

extension Day3 {
    func solve() async throws -> Int {
        switch part {
        case 1:
            var sum = 0
            var processor = LineProcessor()

            for line in try FileReader.lineSequence(path: filePath) {
                let result = processor.process(line)
                logger.log(.debug("buffer \(result)"))
                // let clusters = findDigitClusters(in: line)

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
