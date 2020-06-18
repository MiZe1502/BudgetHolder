import { css } from 'emotion';
import colors from "../../styles/colors";

export { Font716Black } from "../../styles/fonts";
export { SideMainPadding, FlexHorCenter, Overflowed } from "../../styles/positioning";

export const Row = css`
    min-height: 40px;
    padding-top: 8px;
    padding-bottom: 8px;
    align-items: flex-start;

    &:nth-of-type(2n) {
        background-color: ${colors.gray};
    }
`;

export const Cell = css`
    margin-right: 16px;
`;