enum MatchActionType {
    PlaceStone,
    Pass
}

interface Coordinate {
    x: number,
    y: number
}

interface MatchAction {
    action: MatchActionType,
    position: Coordinate | null
}

export default MatchAction;