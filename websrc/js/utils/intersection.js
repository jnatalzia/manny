export function doesOverlap(b1, b2) {
    return (
        b1.x < b2.x + b2.size &&
        b1.x + b1.size > b2.x &&
        b1.y < b2.y + b2.size &&
        b1.y + b1.size > b2.y
    );
}
