import React, {PropTypes} from 'react';
import LayoutRenderer from 'shared/components/LayoutRenderer';
import TimeRangeDropdown from '../../shared/components/TimeRangeDropdown';
import timeRanges from 'hson!../../shared/data/timeRanges.hson';
import {getMappings, getAppsForHosts} from '../apis';
import {fetchLayouts} from 'shared/apis';

export const HostPage = React.createClass({
  propTypes: {
    source: PropTypes.shape({
      links: PropTypes.shape({
        proxy: PropTypes.string.isRequired,
      }).isRequired,
      telegraf: PropTypes.string.isRequired,
    }),
    params: PropTypes.shape({
      hostID: PropTypes.string.isRequired,
    }).isRequired,
    location: PropTypes.shape({
      query: PropTypes.shape({
        app: PropTypes.string,
      }),
    }),
  },

  getInitialState() {
    const fifteenMinutesIndex = 1;

    return {
      layouts: [],
      timeRange: timeRanges[fifteenMinutesIndex],
    };
  },

  componentDidMount() {
    const hosts = {[this.props.params.hostID]: {name: this.props.params.hostID}};
    const {source} = this.props;

    // fetching layouts and mappings can be done at the same time
    fetchLayouts().then(({data: {layouts}}) => {
      getMappings().then(({data: {mappings}}) => {
        getAppsForHosts(source.links.proxy, hosts, mappings, source.telegraf).then((newHosts) => {
          const host = newHosts[this.props.params.hostID];
          const filteredLayouts = layouts.filter((layout) => {
            const focusedApp = this.props.location.query.app;
            if (focusedApp) {
              return layout.app === focusedApp;
            }
            return host.apps && host.apps.includes(layout.app);
          });
          this.setState({layouts: filteredLayouts});
        });
      });
    });
  },

  handleChooseTimeRange({lower}) {
    const timeRange = timeRanges.find((range) => range.queryValue === lower);
    this.setState({timeRange});
  },

  renderLayouts(layouts) {
    const autoRefreshMs = 15000;
    const {timeRange} = this.state;
    const {source} = this.props;

    const allCells = layouts.reduce((all, layout) => {
      return all.concat(layout.cells);
    }, []);

    const adjustedCells = allCells.map((cell, i) => {
      const cellQueries = cell.queries.map((q) => {
        return Object.assign({}, q, {
          text: q.query,
          database: source.telegraf,
        });
      });
      return Object.assign({}, cell, {
        queries: cellQueries,
        x: (i * 4 % 12), // eslint-disable-line no-magic-numbers
        y: 0,
      });
    });

    return (
      <LayoutRenderer
        timeRange={timeRange}
        cells={adjustedCells}
        autoRefreshMs={autoRefreshMs}
        source={source.links.proxy}
        host={this.props.params.hostID}
      />
    );
  },

  render() {
    const hostID = this.props.params.hostID;
    const {layouts, timeRange} = this.state;

    return (
      <div className="host-dashboard hosts-page">
        <div className="chronograf-header hosts-dashboard-header">
          <div className="chronograf-header__container">
            <div className="chronograf-header__left">
              <h1>{hostID}</h1>
            </div>
            <div className="chronograf-header__right">
              <h1>Range:</h1>
              <TimeRangeDropdown onChooseTimeRange={this.handleChooseTimeRange} selected={timeRange.inputValue} />
            </div>
          </div>
        </div>
        <div className="hosts-page-scroll-container">
          <div className="container-fluid hosts-dashboard">
            <div className="row">
              { (layouts.length > 0) ? this.renderLayouts(layouts) : '' }
            </div>
          </div>
        </div>
      </div>
    );
  },
});

export default HostPage;
