import {css} from "emotion";

import { Font312Black } from "../../../styles/fonts";
import { FlexVert } from "../../../styles/positioning";

import colors from "../../../styles/colors";

export const style = {
    Font312Black,
    FlexVert,

    Delimiter: css`
      width: 20px;
      border-bottom: 2px solid ${colors.greenPrimary};
    `,

    SingleLine: css`
        border-left: 2px solid ${colors.greenPrimary};
        padding-left: 8px;
        padding-top: 2px;
    `
}