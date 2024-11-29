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

struct NumberCluster {
    let digits: [Character]
    let startIndex: Int
}

extension Character {
    var isCustomSymbol: Bool {
        return !self.isLetter && !self.isNumber && self != "."
    }
}

func findDigitClusters(in chars: [Character]) -> [NumberCluster] {
    var clusters: [NumberCluster] = []
    var currentDigits: [Character] = []
    var currentStart: Int?

    for (index, char) in chars.enumerated() {
        if char.isNumber {
            if currentStart == nil {
                currentStart = index
            }
            currentDigits.append(char)
        } else if let start = currentStart {
            clusters.append(NumberCluster(digits: currentDigits, startIndex: start))
            currentDigits = []
            currentStart = nil
        }
    }

    if let start = currentStart {
        clusters.append(NumberCluster(digits: currentDigits, startIndex: start))
    }

    return clusters
}

struct Buffer {
    private var first: [Character] = []
    private var second: [Character] = []
    private var third: [Character] = []

    mutating func add(value: [Character]) {
        if full {
            first = second
            second = third
            third = value
        } else {
            if first.count == 0 {
                first = value
            } else if second.count == 0 {
                second = value
            } else {
                third = value
            }
        }
    }

    init() {
    }

    var full: Bool {
        return first.count > 0 && second.count > 0 && third.count > 0
    }
}

extension Buffer {
    func checkFirstLine() -> Int {
        var sum = 0

        for cluster in findDigitClusters(in: first) {
            /// check left
            if cluster.startIndex > 0 {
                if first[cluster.startIndex].isCustomSymbol {
                    if let number = Int(String(cluster.digits)) {
                        sum += number
                        continue
                    }
                }
            }

            /// check right
            if cluster.startIndex + cluster.digits.count + 1 < first.count {
                if first[cluster.startIndex + cluster.digits.count].isCustomSymbol {
                    if let number = Int(String(cluster.digits)) {
                        sum += number
                        continue
                    }
                }
            }

            // check below
            let belowSubset: ArraySlice<Character>
            if cluster.startIndex > 0 {
                belowSubset =
                    second[cluster.startIndex - 1...cluster.startIndex + cluster.digits.count]
            } else {
                if cluster.startIndex + cluster.digits.count + 1 < second.count {
                    belowSubset =
                        second[cluster.startIndex...cluster.startIndex + cluster.digits.count + 1]
                } else {
                    belowSubset =
                        second[cluster.startIndex...cluster.startIndex + cluster.digits.count]
                }
            }

            for char in belowSubset {
                if char.isCustomSymbol {
                    if let number = Int(String(cluster.digits)) {
                        sum += number
                        break
                    }
                }
            }

        }
        return sum
    }

    func checkSecondLine() -> Int {
        var sum = 0

        return sum
    }

    func checkLastLIne() -> Int {
        var sum = 0

        return sum
    }
}

enum CounterPosition {
    case first
    case other
    case last
}

extension Buffer {
    func process(counterPosition: CounterPosition) -> [Int] {
        switch counterPosition {
        case .first:
            break
        case .other:
            break
        case .last:
            break
        }
        return []
    }
}

extension Day3 {
    func solve() async throws -> Int {
        switch part {
        case 1:
            var sum = 0
            var buffer = Buffer()

            for (counter, line) in try FileReader.lineSequence(path: filePath).enumerated() {
                // let result = processor.process(line)
                buffer.add(value: Array(line))
                // some efficiency from switch
                switch (counter, buffer.full) {
                case (..<3, true):
                    // process the first 3 lines
                    sum += buffer.checkFirstLine()
                    sum += buffer.checkSecondLine()
                case (_, true):
                    sum += buffer.checkSecondLine()
                default:
                    break
                }
            }
            // when all is over, check last line
            sum += buffer.checkLastLIne()
            return sum

        case 2:
            var sum = 0

            return sum
        default:
            preconditionFailure("Invalid part")
        }
    }
}
