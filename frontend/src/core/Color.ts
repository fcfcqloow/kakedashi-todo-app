export type Color = string;

export const Color = {
  ERROR   : <Color>'#B00020',
  SUCCESS : <Color>'#28a745',
  WARN    : <Color>'#FFD700',
  BLACK   : <Color>'#000000',
  WHITE   : <Color>'#FFFFFF',
  GRAY    : <Color>'#888888',
};

// type NumStr = '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9';
// type HexStr = NumStr | 'A' | 'B' | 'C' | 'D' | 'E' | 'F';
// type LongHexStr = `${HexStr}${HexStr}${HexStr}`;
// 複雑すぎて許されない
// type ColorStr = `#${LongHexStr}${LongHexStr}` | `#${HexStr}${HexStr}${HexStr}`;

export const getFontColor = (backColor: string): Color => {
  if (![4, 7].includes(backColor.length)) {
    throw new Error('unknown color ' + backColor);
  }

  const isShortColorStr = backColor.length === 4;
  const R = isShortColorStr ? parseInt(backColor.substring(1, 2), 16) * 16 : parseInt(backColor.substring(1, 2), 16) * 16;
  const G = isShortColorStr ? parseInt(backColor.substring(2, 3), 16) * 16 : parseInt(backColor.substring(3, 4), 16) * 16;
  const B = isShortColorStr ? parseInt(backColor.substring(4, 5), 16) * 16 : parseInt(backColor.substring(5, 6), 16) * 16;

  if (R > 130 && G > 150 && B > 150) {
    return Color.BLACK;
  }

  if  (R > 240 || G > 200 || B > 220) {
    return Color.GRAY;
  }

  return Color.WHITE;
};