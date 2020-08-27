
import { css } from 'emotion';
import colors from '../../../styles/colors';

import { FlexHorCenter } from "../../../styles/positioning";
import { Font312Black } from "../../../styles/fonts";

export const style = {
    FlexHorCenter,
    Font312Black,

    SingleButton: css`
        height: 25px;
        width: 35px;
        justify-content: center;
        border-radius: 4px;
        margin-right: 4px;
        padding-bottom: 2px;
    
        &:hover {
            cursor: pointer;
            background-color: ${colors.greenLight};
        }
    `,

    ActiveButton: css`
       background-color: ${colors.greenActive};
    
       &:hover {
        background-color: ${colors.greenActive};
       }
    `
}