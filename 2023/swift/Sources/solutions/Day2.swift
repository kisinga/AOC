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

enum Colors {
    case red(_ value: Int)
    case blue(_ value: Int)
    case green(_ value: Int)
}

struct Game {
    var red: Int?
    var green: Int?
    var blue: Int?

    init(red: Int, green: Int, blue: Int) {
        self.red = red
        self.blue = blue
        self.green = green
    }

    subscript(color: String) -> Int? {
        get {
            switch color.lowercased() {
            case "red":
                return red
            case "green":
                return green
            case "blue":
                return blue
            default:
                return 0
            }
        }
        set {
            switch color.lowercased() {
            case "red":
                red = newValue
            case "green":
                green = newValue
            case "blue":
                blue = newValue
            default:
                break
            }
        }
    }
}

typealias GameConstraint = Game

extension Day2 {
    func solve() async throws -> Int {
        switch part {
        case 1:
            let constraint = Game(red: 12, green: 13, blue: 14)
            var sum = 0
            let pattern = /(\d+)\s+(blue|red|green)/
            for try await line in FileReader.lines(from: filePath) {
                // Let's extract all matches
                var impossible = false
                let id: Int = Int(line.firstMatch(of: /Game (\d+):/)?.1 ?? "0")!
                let matches = line.matches(of: pattern)
                for match in matches {
                    let count = Int(match.1)!
                    let color = match.2

                    if count > constraint[String(color)] ?? 0 {
                        impossible = true
                    }
                }
                if !impossible {
                    logger.log(.debug("Game \(id) possible"))
                    sum += id
                } else {
                    logger.log(.warn("Game \(id) IMPOSSIBLE"))
                }

            }
            return sum

        case 2:
            var powersTotal = 0
            let pattern = /(\d+)\s+(blue|red|green)/
            var biggestValues = Game(red: 0, green: 0, blue: 0)

            for try await line in FileReader.lines(from: filePath) {
                let subsets = line.split(separator: ";")
                for sett in subsets {
                    let matches = sett.matches(of: pattern)
                    for match in matches {
                        let value = Int(match.1)!
                        let color = match.2
                        if biggestValues[String(color)]! < value {
                            biggestValues[String(color)] = value
                        }
                    }    
                }

                let id: Int = Int(line.firstMatch(of: /Game (\d+):/)?.1 ?? "0")!
                logger.log(
                    .debug(
                        "Game \(id) Total \(biggestValues.red ?? 1 ) * \(biggestValues.green ?? 1) * \(biggestValues.blue ?? 1))"
                    ))
                powersTotal +=
                    (biggestValues.red ?? 1) * (biggestValues.green ?? 1)
                    * (biggestValues.blue ?? 1)
                biggestValues = Game(red: 0, green: 0, blue: 0)
            }

            return powersTotal
        default:
            preconditionFailure("Invalid part")
        }
    }
}
