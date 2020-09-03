import { css } from 'emotion';
import colors from '../../../styles/colors';

import { FlexHorCenter } from "../../../styles/positioning";

export const style = {
    FlexHorCenter,

    SingleDot: css`
        height: 5px;
        width: 5px;
        border-radius: 50%;
        background-color: ${colors.greenPrimary};
        margin-right: 4px;
    `
}