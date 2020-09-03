import { css } from 'emotion';
import colors from "../../../styles/colors";

import { Font312Black } from "../../../styles/fonts";
import { FlexHorCenter, Overflowed } from "../../../styles/positioning";

export const style = {
    Font312Black,
    FlexHorCenter,
    Overflowed,

    UrlWrapper: css`
        max-width: 100%;
    `,

    Url: css`
        color: ${colors.orangePrimary};
        font-weight: 500;
        text-decoration: underline;
    
        &:hover {
            color: ${colors.orangeActive};
        }
    
        &:active {
            color: ${colors.orangeLight};
            text-decoration: none;
        }
    
        &:focus {
            outline: none;
        }
    `
}