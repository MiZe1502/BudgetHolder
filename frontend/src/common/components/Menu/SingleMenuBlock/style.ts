import { css } from 'emotion';
import colors from "../../../styles/colors";

import { Font422Black } from "../../../styles/fonts";
import { FlexHorCenter } from "../../../styles/positioning";

export const style = {
    Font422Black,
    FlexHorCenter,

    SingleMenuBlock: css`
        height: 5vh;
        width: auto;
        background-color: ${colors.greenPrimary};
        padding: 8px 12px;
        min-width: 100px;
        justify-content: center;
        transition: background-color 300ms ease;

        &:active,
        &:hover {
            cursor: pointer;
            background-color: ${colors.greenActive};
        }

        &:active {
            background-color: ${colors.greenLight};
        }
    `,

    ActiveBlock: css`
      background-color: ${colors.greenActive};
    `
}
