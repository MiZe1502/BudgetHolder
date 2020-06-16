import { css } from 'emotion';
import colors from "../../styles/colors";

export { Font422Black } from "../../styles/fonts";
export { FlexHorCenter } from "../../styles/positioning";

export const SingleMenuBlock = css`
        height: 5vh;
        width: auto;
        background-color: ${colors.greenPrimary};
        padding-bottom: 8px;
        padding-top: 8px;
        padding-left: 12px;
        padding-right: 12px;
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
`;

export const ActiveBlock = css`
    background-color: ${colors.greenActive};
`;