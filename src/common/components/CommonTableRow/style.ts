import { css } from 'emotion';
import colors from "../../styles/colors";

export { Font716Black } from "../../styles/fonts";
export { SideMainPadding, FlexHorCenter } from "../../styles/positioning";

export const Row = css`
    min-height: 40px;

    &:nth-of-type(2n) {
        background-color: ${colors.gray};
    }
`;